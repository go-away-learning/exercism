// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string            // Name of the resident
	Age     int               // Age of the resident
	Address map[string]string // Address of the resident, with keys like "street", "city", etc.
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	if r.Name == "" {
		return false
	} else if street, ok := r.Address["street"]; ok {
		if street == "" {
			return false
		}
	} else {
		return false
	}
	return true
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	*r = Resident{}
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	var count int
	for _, r := range residents {
		if r.HasRequiredInfo() {
			count++
		}
	}
	return count
}
