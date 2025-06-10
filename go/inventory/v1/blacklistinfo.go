package v1

func (b *BlacklistInfo) Dup() *BlacklistInfo {
	if b == nil {
		return nil
	}

	return &BlacklistInfo{
		SecurityViolationsCount:         b.SecurityViolationsCount,
		ResourceRequirementsUnfulfilled: b.ResourceRequirementsUnfulfilled,
		DeploymentsClosedWhileOffline:   b.DeploymentsClosedWhileOffline,
	}
}
