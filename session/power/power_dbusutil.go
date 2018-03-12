package power

func (v *Manager) setPropLidIsPresent(value bool) (changed bool) {
	if v.LidIsPresent != value {
		v.LidIsPresent = value
		v.emitPropChangedLidIsPresent(value)
		return true
	}
	return false
}

func (v *Manager) emitPropChangedLidIsPresent(value bool) error {
	return v.service.EmitPropertyChanged(v, "LidIsPresent", value)
}

func (v *Manager) setPropOnBattery(value bool) (changed bool) {
	if v.OnBattery != value {
		v.OnBattery = value
		v.emitPropChangedOnBattery(value)
		return true
	}
	return false
}

func (v *Manager) emitPropChangedOnBattery(value bool) error {
	return v.service.EmitPropertyChanged(v, "OnBattery", value)
}

func (v *Manager) setPropWarnLevel(value WarnLevel) (changed bool) {
	if v.WarnLevel != value {
		v.WarnLevel = value
		v.emitPropChangedWarnLevel(value)
		return true
	}
	return false
}

func (v *Manager) emitPropChangedWarnLevel(value WarnLevel) error {
	return v.service.EmitPropertyChanged(v, "WarnLevel", value)
}
