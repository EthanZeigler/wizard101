package parser

import (
	"wizard101/messages"
	"wizard101/messages/quest"
)

var MessageNumberLookUp = [...]messages.ProtocolMessage {
	nil,
	quest.AcceptQuest{},
	quest.CompleteGoal{},
	quest.CompletePersona{},
	quest.CompleteQuest{},
	quest.DeclineQuest{},
	quest.InteractNpc{},
	quest.NpcInfo{},
	quest.PersonaInfo{},
	quest.QuestOffer{},
	quest.RemoveGoal{},
	quest.RemoveQuest{},
	quest.SendGoal{},
	quest.SendNpcOptions{},
	quest.SendQuest{},
}
func buildQuestStruct(packet []byte, structToBuild string) messages.ProtocolMessage {
	switch structToBuild {
	case quest.AcceptQuest{}.String(): return buildAcceptQuest(packet);
	case quest.CompleteGoal{}.String(): return buildCompleteGoal(packet);
	case quest.CompletePersona{}.String(): return buildCompletePersona(packet);
	case quest.CompleteQuest{}.String(): return buildCompleteQuest(packet);
	case quest.DeclineQuest{}.String(): return buildDeclineGoal(packet);
	case quest.InteractNpc{}.String(): return buildInteractNpc(packet);
	case quest.NpcInfo{}.String(): return buildNpcInfo(packet);
	case quest.PersonaInfo{}.String(): return buildPersonaInfo(packet);
	case quest.QuestOffer{}.String(): return buildQuestOffer(packet);
	case quest.RemoveGoal{}.String(): return buildRemoveGoal(packet);
	case quest.RemoveQuest{}.String(): return buildRemoveQuest(packet);
	case quest.SendGoal{}.String(): return buildSendGoal(packet);
	case quest.SendNpcOptions{}.String(): return buildSendNpcOptions(packet);
	case quest.SendQuest{}.String(): return buildSendQuest(packet);
	default: return nil
	}
}

func buildSendQuest(packet []byte) messages.ProtocolMessage {
	packet, questID := ReadGuid(packet)
	packet, quiestNameID := ReadUInt(packet)
	packet, questType := ReadUInt(packet)
	packet, length := ReadUShort(packet)
	packet, questTitle := readString(packet)
	packet, length = ReadUShort(packet)
	packet, questInfo := ReadStringWithLength(packet, length)
	packet, New := ReadUByte(packet)
	packet, length = ReadUShort(packet)
	packet, questMadLibs := ReadStringWithLength(packet, length)
	packet, length = ReadUShort(packet)
	packet, goalData := ReadStringWithLength(packet, length)
	packet, length = ReadUShort(packet)
	packet, rewards := ReadStringWithLength(packet, length)
	packet, length = ReadUShort(packet)
	packet, clientTags := ReadStringWithLength(packet, length)
	packet, noQuestHelper := ReadUByte(packet)
	packet, mainline := ReadUByte(packet)
	packet, skipAutoSelect := ReadUByte(packet)
	packet, petOnly := ReadUByte(packet)

	return quest.SendQuest{QuestID: questID, QuestNameID: quiestNameID, QuestType: questType, QuestTitle: questTitle,
		QuestInfo: questInfo, New: New, QuestMadLibs: questMadLibs,GoalData: goalData,Rewards: rewards,ClientTags: clientTags,
	NoQuestHelper: noQuestHelper, MainLine: mainline, SkipQHAutoSelect: skipAutoSelect, PetOnlyQuest: petOnly}


}

func buildSendNpcOptions(packet []byte) messages.ProtocolMessage {
	return nil
}

func buildSendGoal(packet []byte) messages.ProtocolMessage {
	packet, questID := ReadGuid(packet)
	packet, goalID := ReadGuid(packet)
	packet, goalName := ReadUInt(packet)
	packet, length := ReadUShort(packet)
	packet, goalTitle := ReadStringWithLength(packet,length)
	packet, length = ReadUShort(packet)
	packet, goalLocation := ReadStringWithLength(packet,length)
	packet, length = ReadUShort(packet)
	packet, goalImage := ReadStringWithLength(packet,length)
	packet, length = ReadUShort(packet)
	packet, goalImage2 := ReadStringWithLength(packet,length)
	packet, length = ReadUShort(packet)
	packet, personaName := ReadStringWithLength(packet,length)
	packet,  goalType := ReadUByte(packet)
	packet, goalstatus := ReadUByte(packet)
	packet,  goalcount:= ReadInt(packet)
	packet, goaltotal := ReadInt(packet)
	packet, subscriberGoalTotal := ReadInt(packet)
	packet, useTally := ReadUByte(packet)
	packet, length = ReadUShort(packet)
	packet, tallyString := ReadStringWithLength(packet,length)
	packet, sendType := ReadInt(packet)
	packet, length = ReadUShort(packet)
	packet, goalMadLibs := ReadStringWithLength(packet,length)
	packet, length = ReadUShort(packet)
	packet, clientTags := readString(packet)
	packet, length = ReadUShort(packet)
	packet, patreonIcon := ReadStringWithLength(packet,length)
	packet, noQuestHelper := ReadUByte(packet)
	packet, petOnly := ReadUByte(packet)

	return quest.SendGoal{QuestID: questID, GoalID: goalID, GoalName: goalName, GoalTitle: string(goalTitle),
		GoalLocation: string(goalLocation), GoalImage: string(goalImage), GoalImage2: string(goalImage2), PersonaName: string(personaName),
	GoalType: goalType, GoalStatus: goalstatus, GoalCount: goalcount, GoalTotal: goaltotal, SubscriberGoalTotal: subscriberGoalTotal,
	UseTally: useTally, TallyText: string(tallyString),SendType: sendType, GoalMadLibs: goalMadLibs, ClientTags: []byte(clientTags), PatronIcon: string(patreonIcon),
	NoQuestHelper: noQuestHelper, PetOnlyQuest: petOnly}
}

func buildRemoveQuest(packet []byte) messages.ProtocolMessage {
	return nil
}

func buildRemoveGoal(packet []byte) messages.ProtocolMessage {
	return nil
}

func buildQuestOffer(packet []byte) messages.ProtocolMessage {
	return nil
}

func buildPersonaInfo(packet []byte) messages.ProtocolMessage {
	return nil
}

func buildNpcInfo(packet []byte) messages.ProtocolMessage {
	return nil
}

func buildInteractNpc(packet []byte) messages.ProtocolMessage {
	return nil
}

func buildDeclineGoal(packet []byte) messages.ProtocolMessage {
	return nil
}

func buildCompleteQuest(packet []byte) messages.ProtocolMessage {
	return nil
}

func buildCompletePersona(packet []byte) messages.ProtocolMessage {
	return nil
}

func buildCompleteGoal(packet []byte) messages.ProtocolMessage {
	return nil
}

func buildAcceptQuest(packet []byte) messages.ProtocolMessage {
	return nil
}
