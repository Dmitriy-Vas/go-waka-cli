package go_waka_api

import "encoding/json"

func (w *WakaClient) OrgDashboards(org string) (*OrgDashboards, error) {
	data, err := w.get("users/current/orgs/"+org+"/dashboards", nil)
	if err != nil {
		return nil, err
	}
	var orgDashboards *OrgDashboards
	if err = json.Unmarshal(data, &orgDashboards); err != nil {
		return nil, err
	}
	return orgDashboards, nil
}
