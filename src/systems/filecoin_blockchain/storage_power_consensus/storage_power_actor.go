package storage_power_consensus

import (
	block "github.com/filecoin-project/specs/systems/filecoin_blockchain/struct/block"
	deal "github.com/filecoin-project/specs/systems/filecoin_markets/deal"
	addr "github.com/filecoin-project/specs/systems/filecoin_vm/actor/address"
	util "github.com/filecoin-project/specs/util"
)

const PLEDGE_COLLATERAL_PER_GB = -1 // TODO define

// Actor
func (spa *StoragePowerActor_I) ReportConsensusFault(slasherAddr addr.Address, faultType ConsensusFaultType, proof []block.Block) {
	panic("TODO")

	// Use EC's IsValidConsensusFault method to validate the proof
	// slash block miner's pledge collateral
	// reward slasher
}

func (spa *StoragePowerActor_I) ReportUncommittedPowerFault(cheaterAddr addr.Address, numSectors util.UVarint) {
	panic("TODO")
	// Quite a bit more straightforward since only called by the cron actor (ie publicly verified)

	// slash cheater pledge collateral accordingly based on num sectors faulted
}

func (spa *StoragePowerActor_I) CommitPledgeCollateral(deals []deal.StorageDeal) {

	panic("TODO")
	// check that based on deals (ie sector sizes and num sectors) miner has enough associated balance in the storage miner wallet
	// pledge and associate
}

func (spa *StoragePowerActor_I) DecommitPledgeCollateral(deals []deal.StorageDeal) {
	panic("TODO")
	// must check more than finality post deal expiration
	// return appropriate amount to storage market based on deals
}

// TODO: add Surprise to the chron actor
func (spa *StoragePowerActor_I) Surprise(ticket block.Ticket) []addr.Address {
	surprisedMiners := []addr.Address{}

	// The number of blocks that a challenged miner has to respond
	// TODO: this should be set in.. spa?
	var provingPeriod uint
	// The number of blocks that a challenged miner has to respond
	// TODO: this should be set in.. spa?
	// var postChallengeTime util.UInt

	// The current currBlockHeight
	// TODO: should be found in vm context
	// var currBlockHeight util.UInt

	// The number of miners that are challenged at this block
	numSurprised := uint(len(spa.Miners())) / provingPeriod

	// TODO: seem inefficient but spa.Miners() is now a map from address to power
	minerAddresses := make([]addr.Address, len(spa.Miners()))

	index := 0
	for address := range spa.Miners() {
		minerAddresses[index] = address
		index++
	}

	for i := uint(0); i < numSurprised; i++ {
		// TODO: randomNumber := hash(ticket, i)
		var randomNumber uint
		minerIndex := randomNumber % uint(len(spa.Miners()))
		minerAddress := minerAddresses[minerIndex]
		surprisedMiners = append(surprisedMiners, minerAddress)
		// TODO: minerActor := GetActorFromID(actor).(storage_mining.StorageMinerActor)

		// TODO: send message to StorageMinerActor to update ProvingPeriod
		// TODO: should this update be done after surprisedMiners respond with a successful PoSt?
		// var minerActor storage_mining.StorageMinerActor
		// minerActor.ProvingPeriodEnd_ = currBlockHeight + postChallengeTime
		// SendMessage(sm.ExtendProvingPeriod)
	}

	return surprisedMiners
}
