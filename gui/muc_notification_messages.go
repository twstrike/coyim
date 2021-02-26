package gui

import (
	"github.com/coyim/coyim/i18n"
	"github.com/coyim/coyim/session/muc/data"
)

func getAffiliationUpdateSuccessMessage(nickname string, previousAffiliation, affiliation data.Affiliation) string {
	if affiliation.IsNone() {
		// This is impossible to happen but we need to cover all cases.
		if previousAffiliation.IsNone() {
			return i18n.Localf("%s no longer has a position.", nickname)
		}
		return i18n.Localf("%s is not %s anymore.", nickname, displayNameForAffiliationWithPreposition(previousAffiliation))
	}

	return i18n.Localf("The position of %s was changed to %s.", nickname, displayNameForAffiliation(affiliation))
}

func getMUCNotificationMessageFrom(d interface{}) string {
	switch t := d.(type) {
	case data.AffiliationUpdate:
		return getAffiliationUpdateMessage(t)
	case data.SelfAffiliationUpdate:
		return getSelfAffiliationUpdateMessage(t)
	case data.RoleUpdate:
		return getRoleUpdateMessage(t)
	case data.SelfRoleUpdate:
		return getSelfRoleUpdateMessage(t)
	case data.AffiliationRoleUpdate:
		return getAffiliationRoleUpdateMessage(t)
	case data.SelfAffiliationRoleUpdate:
		return getSelfAffiliationRoleUpdateMessage(t)
	default:
		return ""
	}
}

func getAffiliationUpdateMessage(affiliationUpdate data.AffiliationUpdate) string {
	switch {
	case affiliationUpdate.New.IsNone():
		return getAffiliationRemovedMessage(affiliationUpdate)

	case affiliationUpdate.New.IsBanned():
		return getAffiliationBannedMessage(affiliationUpdate)

	case affiliationUpdate.Previous.IsNone():
		return getAffiliationAddedMessage(affiliationUpdate)

	default:
		return getAffiliationChangedMessage(affiliationUpdate)
	}
}

func getAffiliationRemovedMessage(affiliationUpdate data.AffiliationUpdate) string {
	if affiliationUpdate.Actor == nil {
		if affiliationUpdate.Reason == "" {
			return i18n.Localf("%s is not %s anymore.",
				affiliationUpdate.Nickname,
				displayNameForAffiliationWithPreposition(affiliationUpdate.Previous))
		}

		return i18n.Localf("%s is not %s anymore because: %s.",
			affiliationUpdate.Nickname,
			displayNameForAffiliationWithPreposition(affiliationUpdate.Previous),
			affiliationUpdate.Reason)
	}

	if affiliationUpdate.Reason == "" {
		return i18n.Localf("The %s %s changed the position of %s; %s is not %s anymore.",
			displayNameForAffiliation(affiliationUpdate.Actor.Affiliation),
			affiliationUpdate.Actor.Nickname,
			affiliationUpdate.Nickname,
			affiliationUpdate.Nickname,
			displayNameForAffiliationWithPreposition(affiliationUpdate.Previous))
	}

	return i18n.Localf("The %s %s changed the position of %s; %s is not %s anymore because: %s.",
		displayNameForAffiliation(affiliationUpdate.Actor.Affiliation),
		affiliationUpdate.Actor.Nickname,
		affiliationUpdate.Nickname,
		affiliationUpdate.Nickname,
		displayNameForAffiliationWithPreposition(affiliationUpdate.Previous),
		affiliationUpdate.Reason)
}

func getAffiliationBannedMessage(affiliationUpdate data.AffiliationUpdate) string {
	if affiliationUpdate.Actor == nil {
		if affiliationUpdate.Reason == "" {
			return i18n.Localf("%s was banned from the room.", affiliationUpdate.Nickname)
		}

		return i18n.Localf("%s was banned from the room because: %s.",
			affiliationUpdate.Nickname,
			affiliationUpdate.Reason)
	}

	if affiliationUpdate.Reason == "" {
		return i18n.Localf("The %s %s banned %s from the room.",
			displayNameForAffiliation(affiliationUpdate.Actor.Affiliation),
			affiliationUpdate.Actor.Nickname,
			affiliationUpdate.Nickname)
	}

	return i18n.Localf("The %s %s banned %s from the room because: %s.",
		displayNameForAffiliation(affiliationUpdate.Actor.Affiliation),
		affiliationUpdate.Actor.Nickname,
		affiliationUpdate.Nickname,
		affiliationUpdate.Reason)
}

func getAffiliationAddedMessage(affiliationUpdate data.AffiliationUpdate) string {
	if affiliationUpdate.Actor == nil {
		if affiliationUpdate.Reason == "" {
			return i18n.Localf("%s is now %s.",
				affiliationUpdate.Nickname,
				displayNameForAffiliationWithPreposition(affiliationUpdate.New))
		}

		return i18n.Localf("%s is now %s because: %s.",
			affiliationUpdate.Nickname,
			displayNameForAffiliationWithPreposition(affiliationUpdate.New),
			affiliationUpdate.Reason)
	}

	if affiliationUpdate.Reason == "" {
		return i18n.Localf("The %s %s changed the position of %s; %s is now %s.",
			displayNameForAffiliation(affiliationUpdate.Actor.Affiliation),
			affiliationUpdate.Actor.Nickname,
			affiliationUpdate.Nickname,
			affiliationUpdate.Nickname,
			displayNameForAffiliationWithPreposition(affiliationUpdate.New))
	}

	return i18n.Localf("The %s %s changed the position of %s; %s is now %s because: %s.",
		displayNameForAffiliation(affiliationUpdate.Actor.Affiliation),
		affiliationUpdate.Actor.Nickname,
		affiliationUpdate.Nickname,
		affiliationUpdate.Nickname,
		displayNameForAffiliationWithPreposition(affiliationUpdate.New),
		affiliationUpdate.Reason)
}

func getAffiliationChangedMessage(affiliationUpdate data.AffiliationUpdate) string {
	if affiliationUpdate.Actor == nil {
		if affiliationUpdate.Reason == "" {
			return i18n.Localf("The position of %s was changed from %s to %s.",
				affiliationUpdate.Nickname,
				displayNameForAffiliation(affiliationUpdate.Previous),
				displayNameForAffiliation(affiliationUpdate.New))
		}

		return i18n.Localf("The position of %s was changed from %s to %s because: %s.",
			affiliationUpdate.Nickname,
			displayNameForAffiliation(affiliationUpdate.Previous),
			displayNameForAffiliation(affiliationUpdate.New),
			affiliationUpdate.Reason)
	}

	if affiliationUpdate.Reason == "" {
		return i18n.Localf("The %s %s changed the position of %s from %s to %s.",
			displayNameForAffiliation(affiliationUpdate.Actor.Affiliation),
			affiliationUpdate.Actor.Nickname,
			affiliationUpdate.Nickname,
			displayNameForAffiliation(affiliationUpdate.Previous),
			displayNameForAffiliation(affiliationUpdate.New))
	}

	return i18n.Localf("The %s %s changed the position of %s from %s to %s because: %s.",
		displayNameForAffiliation(affiliationUpdate.Actor.Affiliation),
		affiliationUpdate.Actor.Nickname,
		affiliationUpdate.Nickname,
		displayNameForAffiliation(affiliationUpdate.Previous),
		displayNameForAffiliation(affiliationUpdate.New),
		affiliationUpdate.Reason)
}

func getRoleUpdateMessage(roleUpdate data.RoleUpdate) string {
	if roleUpdate.Actor == nil {
		if roleUpdate.Reason == "" {
			return i18n.Localf("The role of %s was changed from %s to %s.",
				roleUpdate.Nickname,
				displayNameForRole(roleUpdate.Previous),
				displayNameForRole(roleUpdate.New))
		}

		return i18n.Localf("The role of %s was changed from %s to %s because: %s.",
			roleUpdate.Nickname,
			displayNameForRole(roleUpdate.Previous),
			displayNameForRole(roleUpdate.New),
			roleUpdate.Reason)
	}

	if roleUpdate.Reason == "" {
		return i18n.Localf("The %s %s changed the role of %s from %s to %s.",
			displayNameForAffiliation(roleUpdate.Actor.Affiliation),
			roleUpdate.Actor.Nickname,
			roleUpdate.Nickname,
			displayNameForRole(roleUpdate.Previous),
			displayNameForRole(roleUpdate.New))
	}

	return i18n.Localf("The %s %s changed the role of %s from %s to %s because: %s.",
		displayNameForAffiliation(roleUpdate.Actor.Affiliation),
		roleUpdate.Actor.Nickname,
		roleUpdate.Nickname,
		displayNameForRole(roleUpdate.Previous),
		displayNameForRole(roleUpdate.New),
		roleUpdate.Reason)
}

func getSelfRoleUpdateMessage(selfRoleUpdate data.SelfRoleUpdate) string {
	if selfRoleUpdate.Actor == nil {
		if selfRoleUpdate.Reason == "" {
			return i18n.Localf("Your role was changed from %s to %s.",
				displayNameForRole(selfRoleUpdate.Previous),
				displayNameForRole(selfRoleUpdate.New))
		}

		return i18n.Localf("Your role was changed from %s to %s because: %s.",
			displayNameForRole(selfRoleUpdate.Previous),
			displayNameForRole(selfRoleUpdate.New),
			selfRoleUpdate.Reason)
	}

	if selfRoleUpdate.Reason == "" {
		return i18n.Localf("The %s %s changed your role from %s to %s.",
			displayNameForAffiliation(selfRoleUpdate.Actor.Affiliation),
			selfRoleUpdate.Actor.Nickname,
			displayNameForRole(selfRoleUpdate.Previous),
			displayNameForRole(selfRoleUpdate.New))
	}

	return i18n.Localf("The %s %s changed your role from %s to %s because: %s.",
		displayNameForAffiliation(selfRoleUpdate.Actor.Affiliation),
		selfRoleUpdate.Actor.Nickname,
		displayNameForRole(selfRoleUpdate.Previous),
		displayNameForRole(selfRoleUpdate.New),
		selfRoleUpdate.Reason)
}

func getAffiliationRoleUpdateMessage(affiliationRoleUpdate data.AffiliationRoleUpdate) string {
	switch {
	case affiliationRoleUpdate.NewAffiliation.IsNone() &&
		affiliationRoleUpdate.PreviousAffiliation.IsDifferentFrom(affiliationRoleUpdate.NewAffiliation):
		return getAffiliationRoleUpateForAffiliationRemoved(affiliationRoleUpdate)
	case affiliationRoleUpdate.PreviousAffiliation.IsNone() &&
		affiliationRoleUpdate.NewAffiliation.IsDifferentFrom(affiliationRoleUpdate.PreviousAffiliation):
		return getAffiliationRoleUpdateForAffiliationAdded(affiliationRoleUpdate)
	case affiliationRoleUpdate.NewAffiliation.IsDifferentFrom(affiliationRoleUpdate.PreviousAffiliation):
		return getAffiliationRoleUpdateForAffiliationUpdated(affiliationRoleUpdate)
	default:
		return getAffiliationRoleUpdateForUnexpectedSituation(affiliationRoleUpdate)
	}
}

func getAffiliationRoleUpateForAffiliationRemoved(affiliationRoleUpdate data.AffiliationRoleUpdate) string {
	if affiliationRoleUpdate.Actor == nil {
		if affiliationRoleUpdate.Reason == "" {
			return i18n.Localf("%s is not %s anymore. As a result, their role was changed from %s to %s.",
				affiliationRoleUpdate.Nickname,
				displayNameForAffiliationWithPreposition(affiliationRoleUpdate.PreviousAffiliation),
				displayNameForRole(affiliationRoleUpdate.PreviousRole),
				displayNameForRole(affiliationRoleUpdate.NewRole))
		}

		return i18n.Localf("%s is not %s anymore. As a result, their role was changed from %s to %s. The reason given was: %s.",
			affiliationRoleUpdate.Nickname,
			displayNameForAffiliationWithPreposition(affiliationRoleUpdate.PreviousAffiliation),
			displayNameForRole(affiliationRoleUpdate.PreviousRole),
			displayNameForRole(affiliationRoleUpdate.NewRole),
			affiliationRoleUpdate.Reason)
	}

	if affiliationRoleUpdate.Reason == "" {
		return i18n.Localf("The %s %s changed the position of %s; %s is not an %s anymore. As a result, their role was changed from %s to %s.",
			displayNameForAffiliation(affiliationRoleUpdate.Actor.Affiliation),
			affiliationRoleUpdate.Actor.Nickname,
			affiliationRoleUpdate.Nickname,
			affiliationRoleUpdate.Nickname,
			displayNameForAffiliation(affiliationRoleUpdate.PreviousAffiliation),
			displayNameForRole(affiliationRoleUpdate.PreviousRole),
			displayNameForRole(affiliationRoleUpdate.NewRole))
	}

	return i18n.Localf("The %s %s changed the position of %s; %s is not an %s anymore. "+
		"As a result, their role was changed from %s to %s. The reason given was: %s.",
		displayNameForAffiliation(affiliationRoleUpdate.Actor.Affiliation),
		affiliationRoleUpdate.Actor.Nickname,
		affiliationRoleUpdate.Nickname,
		affiliationRoleUpdate.Nickname,
		displayNameForAffiliation(affiliationRoleUpdate.PreviousAffiliation),
		displayNameForRole(affiliationRoleUpdate.PreviousRole),
		displayNameForRole(affiliationRoleUpdate.NewRole),
		affiliationRoleUpdate.Reason)
}

func getAffiliationRoleUpdateForAffiliationAdded(affiliationRoleUpdate data.AffiliationRoleUpdate) string {
	if affiliationRoleUpdate.Actor == nil {
		if affiliationRoleUpdate.Reason == "" {
			return i18n.Localf("The position of %s was changed to %s. As a result, their role was changed from %s to %s.",
				affiliationRoleUpdate.Nickname,
				displayNameForAffiliation(affiliationRoleUpdate.NewAffiliation),
				displayNameForRole(affiliationRoleUpdate.PreviousRole),
				displayNameForRole(affiliationRoleUpdate.NewRole))
		}

		return i18n.Localf("The position of %s was changed to %s. "+
			"As a result, their role was changed from %s to %s. The reason given was: %s.",
			affiliationRoleUpdate.Nickname,
			displayNameForAffiliation(affiliationRoleUpdate.NewAffiliation),
			displayNameForRole(affiliationRoleUpdate.PreviousRole),
			displayNameForRole(affiliationRoleUpdate.NewRole),
			affiliationRoleUpdate.Reason)
	}

	if affiliationRoleUpdate.Reason == "" {
		return i18n.Localf("The %s %s changed the position of %s to %s. "+
			"As a result, their role was changed from visitor to moderator.",
			displayNameForAffiliation(affiliationRoleUpdate.Actor.Affiliation),
			affiliationRoleUpdate.Actor.Nickname,
			affiliationRoleUpdate.Nickname,
			displayNameForAffiliation(affiliationRoleUpdate.NewAffiliation))
	}

	return i18n.Localf("The %s %s changed the position of %s to %s. "+
		"As a result, their role was changed from %s to %s. The reason given was: %s.",
		displayNameForAffiliation(affiliationRoleUpdate.Actor.Affiliation),
		affiliationRoleUpdate.Actor.Nickname,
		affiliationRoleUpdate.Nickname,
		displayNameForAffiliation(affiliationRoleUpdate.NewAffiliation),
		displayNameForRole(affiliationRoleUpdate.PreviousRole),
		displayNameForRole(affiliationRoleUpdate.NewRole),
		affiliationRoleUpdate.Reason)
}

func getAffiliationRoleUpdateForAffiliationUpdated(affiliationRoleUpdate data.AffiliationRoleUpdate) string {
	if affiliationRoleUpdate.Actor == nil {
		if affiliationRoleUpdate.Reason == "" {
			return i18n.Localf("The position of %s was changed from %s to %s. As a result, their role was changed from %s to %s.",
				affiliationRoleUpdate.Nickname,
				displayNameForAffiliation(affiliationRoleUpdate.PreviousAffiliation),
				displayNameForAffiliation(affiliationRoleUpdate.NewAffiliation),
				displayNameForRole(affiliationRoleUpdate.PreviousRole),
				displayNameForRole(affiliationRoleUpdate.NewRole))
		}

		return i18n.Localf("The position of %s was changed from %s to %s. "+
			"As a result, their role was changed from %s to %s. The reason given was: %s.",
			affiliationRoleUpdate.Nickname,
			displayNameForAffiliation(affiliationRoleUpdate.PreviousAffiliation),
			displayNameForAffiliation(affiliationRoleUpdate.NewAffiliation),
			displayNameForRole(affiliationRoleUpdate.PreviousRole),
			displayNameForRole(affiliationRoleUpdate.NewRole),
			affiliationRoleUpdate.Reason)
	}

	if affiliationRoleUpdate.Reason == "" {
		return i18n.Localf("The %s %s changed the position of %s from %s to %s. "+
			"As a result, their role was changed from visitor to moderator.",
			displayNameForAffiliation(affiliationRoleUpdate.Actor.Affiliation),
			affiliationRoleUpdate.Actor.Nickname,
			affiliationRoleUpdate.Nickname,
			displayNameForAffiliation(affiliationRoleUpdate.PreviousAffiliation),
			displayNameForAffiliation(affiliationRoleUpdate.NewAffiliation))
	}

	return i18n.Localf("The %s %s changed the position of %s from %s to %s. "+
		"As a result, their role was changed from %s to %s. The reason given was: %s.",
		displayNameForAffiliation(affiliationRoleUpdate.Actor.Affiliation),
		affiliationRoleUpdate.Actor.Nickname,
		affiliationRoleUpdate.Nickname,
		displayNameForAffiliation(affiliationRoleUpdate.PreviousAffiliation),
		displayNameForAffiliation(affiliationRoleUpdate.NewAffiliation),
		displayNameForRole(affiliationRoleUpdate.PreviousRole),
		displayNameForRole(affiliationRoleUpdate.NewRole),
		affiliationRoleUpdate.Reason)
}

func getAffiliationRoleUpdateForUnexpectedSituation(affiliationRoleUpdate data.AffiliationRoleUpdate) string {
	if affiliationRoleUpdate.Actor == nil {
		if affiliationRoleUpdate.Reason == "" {
			return i18n.Localf("The position and the role of %s were changed.",
				affiliationRoleUpdate.Nickname)
		}

		return i18n.Localf("The position and the role of %s were changed because: %s.",
			affiliationRoleUpdate.Nickname,
			affiliationRoleUpdate.Reason)
	}

	if affiliationRoleUpdate.Reason == "" {
		return i18n.Localf("The %s %s changed the position of %s. As a result, their role was also changed.",
			displayNameForAffiliation(affiliationRoleUpdate.Actor.Affiliation),
			affiliationRoleUpdate.Actor.Nickname,
			affiliationRoleUpdate.Nickname)
	}

	return i18n.Localf("The %s %s changed the position of %s. "+
		"As a result, their role was also changed. The reason given was: %s.",
		displayNameForAffiliation(affiliationRoleUpdate.Actor.Affiliation),
		affiliationRoleUpdate.Actor.Nickname,
		affiliationRoleUpdate.Nickname,
		affiliationRoleUpdate.Reason)
}

func getSelfAffiliationRoleUpdateMessage(selfAffiliationRoleUpdate data.SelfAffiliationRoleUpdate) string {
	switch {
	case selfAffiliationRoleUpdate.NewAffiliation.IsNone() &&
		selfAffiliationRoleUpdate.PreviousAffiliation.IsDifferentFrom(selfAffiliationRoleUpdate.NewAffiliation):
		return getSelfAffiliationRoleUpateForAffiliationRemoved(selfAffiliationRoleUpdate)
	case selfAffiliationRoleUpdate.PreviousAffiliation.IsNone() &&
		selfAffiliationRoleUpdate.NewAffiliation.IsDifferentFrom(selfAffiliationRoleUpdate.PreviousAffiliation):
		return getSelfAffiliationRoleUpdateForAffiliationAdded(selfAffiliationRoleUpdate)
	default:
		return getSelfAffiliationRoleUpdateForUnexpectedSituation(selfAffiliationRoleUpdate)
	}
}

func getSelfAffiliationRoleUpateForAffiliationRemoved(selfAffiliationRoleUpdate data.SelfAffiliationRoleUpdate) string {
	if selfAffiliationRoleUpdate.Actor == nil {
		if selfAffiliationRoleUpdate.Reason == "" {
			return i18n.Localf("You are not %s anymore. As a result, your role was changed from %s to %s.",
				displayNameForAffiliationWithPreposition(selfAffiliationRoleUpdate.PreviousAffiliation),
				displayNameForRole(selfAffiliationRoleUpdate.PreviousRole),
				displayNameForRole(selfAffiliationRoleUpdate.NewRole))
		}

		return i18n.Localf("You are not %s anymore. As a result, your role was changed from %s to %s. The reason given was: %s.",
			displayNameForAffiliationWithPreposition(selfAffiliationRoleUpdate.PreviousAffiliation),
			displayNameForRole(selfAffiliationRoleUpdate.PreviousRole),
			displayNameForRole(selfAffiliationRoleUpdate.NewRole),
			selfAffiliationRoleUpdate.Reason)
	}

	if selfAffiliationRoleUpdate.Reason == "" {
		return i18n.Localf("The %s %s changed your position; you are not %s anymore. As a result, your role was changed from %s to %s.",
			displayNameForAffiliation(selfAffiliationRoleUpdate.Actor.Affiliation),
			selfAffiliationRoleUpdate.Actor.Nickname,
			displayNameForAffiliationWithPreposition(selfAffiliationRoleUpdate.PreviousAffiliation),
			displayNameForRole(selfAffiliationRoleUpdate.PreviousRole),
			displayNameForRole(selfAffiliationRoleUpdate.NewRole))
	}

	return i18n.Localf("The %s %s changed your position; you are not %s anymore. As a result, your role was changed from %s to %s. The reason given was: %s.",
		displayNameForAffiliation(selfAffiliationRoleUpdate.Actor.Affiliation),
		selfAffiliationRoleUpdate.Actor.Nickname,
		displayNameForAffiliationWithPreposition(selfAffiliationRoleUpdate.PreviousAffiliation),
		displayNameForRole(selfAffiliationRoleUpdate.PreviousRole),
		displayNameForRole(selfAffiliationRoleUpdate.NewRole),
		selfAffiliationRoleUpdate.Reason)
}

func getSelfAffiliationRoleUpdateForAffiliationAdded(selfAffiliationRoleUpdate data.SelfAffiliationRoleUpdate) string {
	if selfAffiliationRoleUpdate.Actor == nil {
		if selfAffiliationRoleUpdate.Reason == "" {
			return i18n.Localf("Your position was changed to %s. As a result, your role was changed from %s to %s.",
				displayNameForAffiliation(selfAffiliationRoleUpdate.NewAffiliation),
				displayNameForRole(selfAffiliationRoleUpdate.PreviousRole),
				displayNameForRole(selfAffiliationRoleUpdate.NewRole))
		}

		return i18n.Localf("Your position was changed to %s. As a result, your role was changed from %s to %s. The reason given was: %s.",
			displayNameForAffiliation(selfAffiliationRoleUpdate.NewAffiliation),
			displayNameForRole(selfAffiliationRoleUpdate.PreviousRole),
			displayNameForRole(selfAffiliationRoleUpdate.NewRole),
			selfAffiliationRoleUpdate.Reason)
	}

	if selfAffiliationRoleUpdate.Reason == "" {
		return i18n.Localf("The %s %s changed your position to %s. As a result, your role was changed from %s to %s.",
			displayNameForAffiliation(selfAffiliationRoleUpdate.Actor.Affiliation),
			selfAffiliationRoleUpdate.Actor.Nickname,
			displayNameForAffiliation(selfAffiliationRoleUpdate.NewAffiliation),
			displayNameForRole(selfAffiliationRoleUpdate.PreviousRole),
			displayNameForRole(selfAffiliationRoleUpdate.NewRole))
	}

	return i18n.Localf("The %s %s changed your position to %s. As a result, your role was changed from %s to %s. The reason given was: %s.",
		displayNameForAffiliation(selfAffiliationRoleUpdate.Actor.Affiliation),
		selfAffiliationRoleUpdate.Actor.Nickname,
		displayNameForAffiliation(selfAffiliationRoleUpdate.NewAffiliation),
		displayNameForRole(selfAffiliationRoleUpdate.PreviousRole),
		displayNameForRole(selfAffiliationRoleUpdate.NewRole),
		selfAffiliationRoleUpdate.Reason)
}

func getSelfAffiliationRoleUpdateForUnexpectedSituation(selfAffiliationRoleUpdate data.SelfAffiliationRoleUpdate) string {
	if selfAffiliationRoleUpdate.Actor == nil {
		if selfAffiliationRoleUpdate.Reason == "" {
			return i18n.Localf("Your position and the role were changed.")
		}

		return i18n.Localf("Your position and the role were changed because: %s.",
			selfAffiliationRoleUpdate.Reason)
	}

	if selfAffiliationRoleUpdate.Reason == "" {
		return i18n.Localf("The %s %s changed your position. As a result, your role was also changed.",
			displayNameForAffiliation(selfAffiliationRoleUpdate.Actor.Affiliation),
			selfAffiliationRoleUpdate.Actor.Nickname)
	}

	return i18n.Localf("The %s %s changed your position. "+
		"As a result, your role was also changed. The reason given was: %s.",
		displayNameForAffiliation(selfAffiliationRoleUpdate.Actor.Affiliation),
		selfAffiliationRoleUpdate.Actor.Nickname,
		selfAffiliationRoleUpdate.Reason)
}

func getSelfAffiliationUpdateMessage(selfAffiliationUpdate data.SelfAffiliationUpdate) string {
	switch {
	case selfAffiliationUpdate.New.IsNone():
		return getSelfAffiliationRemovedMessage(selfAffiliationUpdate)
	case selfAffiliationUpdate.New.IsBanned():
		return getSelfAffiliationBannedMessage(selfAffiliationUpdate)
	case selfAffiliationUpdate.Previous.IsNone():
		return getSelfAffiliationAddedMessage(selfAffiliationUpdate)
	default:
		return getSelfAffiliationChangedMessage(selfAffiliationUpdate)
	}
}

func getSelfAffiliationRemovedMessage(selfAffiliationUpdate data.SelfAffiliationUpdate) string {
	if selfAffiliationUpdate.Actor == nil {
		if selfAffiliationUpdate.Reason == "" {
			return i18n.Localf("You are not %s anymore.",
				displayNameForAffiliationWithPreposition(selfAffiliationUpdate.Previous))
		}

		return i18n.Localf("You are not %s anymore. The reason given was: %s.",
			displayNameForAffiliationWithPreposition(selfAffiliationUpdate.Previous),
			selfAffiliationUpdate.Reason)
	}

	if selfAffiliationUpdate.Reason == "" {
		return i18n.Localf("The %s %s changed your position; you are not %s anymore.",
			displayNameForAffiliation(selfAffiliationUpdate.Actor.Affiliation),
			selfAffiliationUpdate.Actor.Nickname,
			displayNameForAffiliationWithPreposition(selfAffiliationUpdate.Previous))
	}

	return i18n.Localf("The %s %s changed your position; you are not %s anymore. The reason given was: %s.",
		displayNameForAffiliation(selfAffiliationUpdate.Actor.Affiliation),
		selfAffiliationUpdate.Actor.Nickname,
		displayNameForAffiliationWithPreposition(selfAffiliationUpdate.Previous),
		selfAffiliationUpdate.Reason)
}

func getSelfAffiliationAddedMessage(selfAffiliationUpdate data.SelfAffiliationUpdate) string {
	if selfAffiliationUpdate.Actor == nil {
		if selfAffiliationUpdate.Reason == "" {
			return i18n.Localf("You are now %s.", displayNameForAffiliationWithPreposition(selfAffiliationUpdate.New))
		}

		return i18n.Localf("You are now %s. The reason given was: %s.",
			displayNameForAffiliationWithPreposition(selfAffiliationUpdate.New),
			selfAffiliationUpdate.Reason)
	}

	if selfAffiliationUpdate.Reason == "" {
		return i18n.Localf("The %s %s changed your position; you are now %s.",
			displayNameForAffiliation(selfAffiliationUpdate.Actor.Affiliation),
			selfAffiliationUpdate.Actor.Nickname,
			displayNameForAffiliationWithPreposition(selfAffiliationUpdate.New))
	}

	return i18n.Localf("The %s %s changed your position; you are now %s. The reason given was: %s.",
		displayNameForAffiliation(selfAffiliationUpdate.Actor.Affiliation),
		selfAffiliationUpdate.Actor.Nickname,
		displayNameForAffiliationWithPreposition(selfAffiliationUpdate.New),
		selfAffiliationUpdate.Reason)
}

func getSelfAffiliationChangedMessage(selfAffiliationUpdate data.SelfAffiliationUpdate) string {
	if selfAffiliationUpdate.Actor == nil {
		if selfAffiliationUpdate.Reason == "" {
			return i18n.Localf("Your position was changed from %s to %s.",
				displayNameForAffiliation(selfAffiliationUpdate.Previous),
				displayNameForAffiliation(selfAffiliationUpdate.New))
		}

		return i18n.Localf("Your position was changed from %s to %s because: %s.",
			displayNameForAffiliation(selfAffiliationUpdate.Previous),
			displayNameForAffiliation(selfAffiliationUpdate.New),
			selfAffiliationUpdate.Reason)
	}

	if selfAffiliationUpdate.Reason == "" {
		return i18n.Localf("The %s %s changed your position from %s to %s.",
			displayNameForAffiliation(selfAffiliationUpdate.Actor.Affiliation),
			selfAffiliationUpdate.Actor.Nickname,
			displayNameForAffiliation(selfAffiliationUpdate.Previous),
			displayNameForAffiliation(selfAffiliationUpdate.New))
	}

	return i18n.Localf("The %s %s changed your position from %s to %s because: %s.",
		displayNameForAffiliation(selfAffiliationUpdate.Actor.Affiliation),
		selfAffiliationUpdate.Actor.Nickname,
		displayNameForAffiliation(selfAffiliationUpdate.Previous),
		displayNameForAffiliation(selfAffiliationUpdate.New),
		selfAffiliationUpdate.Reason)
}

func getSelfAffiliationBannedMessage(selfAffiliationUpdate data.SelfAffiliationUpdate) string {
	if selfAffiliationUpdate.Actor == nil {
		if selfAffiliationUpdate.Reason == "" {
			return i18n.Localf("You has been banned from the room.")
		}

		return i18n.Localf("You has been banned from the room. The reason given was: %s.", selfAffiliationUpdate.Reason)
	}

	if selfAffiliationUpdate.Reason == "" {
		return i18n.Localf("The %s %s banned you from the room.",
			displayNameForAffiliation(selfAffiliationUpdate.Actor.Affiliation),
			selfAffiliationUpdate.Actor.Nickname)
	}

	return i18n.Localf("The %s %s banned you from the room. The reason given was: %s.",
		displayNameForAffiliation(selfAffiliationUpdate.Actor.Affiliation),
		selfAffiliationUpdate.Actor.Nickname,
		selfAffiliationUpdate.Reason)
}
