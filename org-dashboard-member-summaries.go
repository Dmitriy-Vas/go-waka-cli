package go_waka_api

import "encoding/json"

func (w *WakaClient) OrgDashboardMemberSummaries(org string, dashboard string, member string, start *Data, end *Data) (*OrgDashboardMemberSummaries, error) {
	query := make(map[string]string)
	query["start"] = start.format()
	query["end"] = end.format()
	data, err := w.get("users/current/orgs/"+org+"/dashboards/"+dashboard+"/members/"+member+"/summaries", query)
	if err != nil {
		return nil, err
	}
	var orgDashboardMemberSummaries *OrgDashboardMemberSummaries
	if err = json.Unmarshal(data, &orgDashboardMemberSummaries); err != nil {
		return nil, err
	}
	return orgDashboardMemberSummaries, nil
}
