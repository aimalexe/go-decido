package decision

// RebuildCounters restores internal IDs after unmarshalling a Decision.
// The counters are intentionally unexported and therefore are not stored in JSON.
func (d *Decision) RebuildCounters() {
	var maxCriterionID int
	var maxAlternativeID int

	for _, criterion := range d.Criteria {
		if criterion.ID > maxCriterionID {
			maxCriterionID = criterion.ID
		}
	}

	for _, alternative := range d.Alternatives {
		if alternative.ID > maxAlternativeID {
			maxAlternativeID = alternative.ID
		}
	}

	d.nextCriterionID = maxCriterionID + 1
	d.nextAlternativeID = maxAlternativeID + 1
}
