package api

import "encoding/json"

func (w *WakaClient) OrgDashboardMembers(org string, dashboard string) (*OrgDashboardMembers, error) {
	data, err := w.get("users/current/orgs/"+org+"/dashboards/"+dashboard+"/members", nil)
	if err != nil {
		return nil, err
	}
	var orgDashboardMembers *OrgDashboardMembers
	if err = json.Unmarshal(data, &orgDashboardMembers); err != nil {
		return nil, err
	}
	return orgDashboardMembers, nil
}
