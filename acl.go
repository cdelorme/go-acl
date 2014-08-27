package acl

type Acl struct {
    Groups map[string][]string
    Actions []string
}

func (acl *Acl) build() {
    if acl.Groups == nil {
        acl.Groups = make(map[string][]string)
    }
}

func (acl *Acl) addGroup(group string) {
    acl.build()
    if _, ok := acl.Groups[group]; !ok {
        acl.Groups[group] = make([]string, 0)
    }
}

func (acl *Acl) hasAction(action string) bool {
    for _, a := range acl.Actions {
        if a == action {
            return true
        }
    }
    return false
}

func (acl *Acl) addActions(actions ...string) {
    for _, action := range actions {
        if !acl.hasAction(action) {
            acl.Actions = append(acl.Actions, action)
        }
    }
}

func (acl *Acl) Grant(group string, actions ...string) {
    acl.addActions(actions...)
    acl.addGroup(group)
    acl.Groups[group] = append(acl.Groups[group], actions...)
}

func (acl *Acl) Revoke(group string, actions ...string) {
    acl.addGroup(group)
    for i, a := range acl.Groups[group] {
        for _, action := range actions {
            if a == action {
                acl.Groups[group] = append(acl.Groups[group][:i], acl.Groups[group][i+1:]...)
            }
        }
    }
}

func (acl *Acl) Can(group string, action string) bool {
    acl.addGroup(group)
    for _, a := range acl.Groups[group] {
        if a == action {
            return true
        }
    }
    return false
}

func (acl *Acl) Not(group string, action string) bool {
    return !acl.Can(group, action)
}

func (acl *Acl) Clear() {
    acl.Groups = nil
}

