package pre2p

// TODO(team): This is a a temporary parallel to the real `p2p` module.
// It should be removed once the real `p2p` module is ready but is meant
// to be a "real" replacement for now.

import (
	"log"

	"github.com/pokt-network/pocket/p2p/pre2p/raintree"
	"github.com/pokt-network/pocket/p2p/pre2p/stdnetwork"
	typesPre2P "github.com/pokt-network/pocket/p2p/pre2p/types"

	"github.com/pokt-network/pocket/shared/config"
	cryptoPocket "github.com/pokt-network/pocket/shared/crypto"
	"github.com/pokt-network/pocket/shared/modules"
	"github.com/pokt-network/pocket/shared/types"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

var _ modules.P2PModule = &p2pModule{}

type p2pModule struct {
	bus       modules.Bus
	p2pConfig *config.Pre2PConfig

	listener typesPre2P.Transport
	address  cryptoPocket.Address

	network typesPre2P.Network
}

func Create(cfg *config.Config) (m modules.P2PModule, err error) {
	log.Println("Creating network module")

	l, err := CreateListener(cfg.Pre2P)
	if err != nil {
		return nil, err
	}

	m = &p2pModule{
		p2pConfig: cfg.Pre2P,

		listener: l,
		address:  cfg.PrivateKey.Address(),

		network: nil,
	}

	return m, nil
}

func (m *p2pModule) SetBus(bus modules.Bus) {
	m.bus = bus
}

func (m *p2pModule) GetBus() modules.Bus {
	if m.bus == nil {
		log.Fatalf("PocketBus is not initialized")
	}
	return m.bus
}

func (m *p2pModule) Start() error {
	log.Println("Starting network module")

	addrBook, err := ValidatorMapToAddrBook(m.p2pConfig, m.bus.GetConsensusModule().ValidatorMap())
	if err != nil {
		return err
	}

	if m.p2pConfig.UseRainTree {
		m.network = raintree.NewRainTreeNetwork(m.address, addrBook)
	} else {
		m.network = stdnetwork.NewNetwork(addrBook)
	}

	go func() {
		for {
			data, err := m.listener.Read()
			if err != nil {
				log.Println("Error reading data from connection: ", err)
				continue
			}
			go m.handleNetworkMessage(data)
		}
	}()

	return nil
}

func (m *p2pModule) Stop() error {
	log.Println("Stopping network module")
	if err := m.listener.Close(); err != nil {
		return err
	}
	return nil
}

func (m *p2pModule) Broadcast(msg *anypb.Any, topic types.PocketTopic) error {
	c := &types.PocketEvent{
		Topic: topic,
		Data:  msg,
	}
	data, err := proto.Marshal(c)
	if err != nil {
		return err
	}
	log.Println("broadcasting message to network")
	return m.network.NetworkBroadcast(data)
}

func (m *p2pModule) Send(addr cryptoPocket.Address, msg *anypb.Any, topic types.PocketTopic) error {
	c := &types.PocketEvent{
		Topic: topic,
		Data:  msg,
	}
	data, err := proto.Marshal(c)
	if err != nil {
		return err
	}

	return m.network.NetworkSend(data, addr)
}

func (m *p2pModule) handleNetworkMessage(networkMsgData []byte) {
	appMsgData, err := m.network.HandleNetworkData(networkMsgData)
	if err != nil {
		log.Println("Error handling raw data: ", err)
		return
	}

	// There was no error, but we don't need to forward this to the app-specific bus.
	// For example, the message has already been handled by the application.
	if appMsgData == nil {
		// log.Println("[DEBUG] No app-specific message to forward from the network")
		return
	}

	networkMessage := types.PocketEvent{}
	if err := proto.Unmarshal(appMsgData, &networkMessage); err != nil {
		log.Println("Error decoding network message: ", err)
		return
	}

	event := types.PocketEvent{
		Topic: networkMessage.Topic,
		Data:  networkMessage.Data,
	}

	m.GetBus().PublishEventToBus(&event)
}
