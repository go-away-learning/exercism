// Package gross provides functions for managing customer bills and unit measurements at the Gross Store.
package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if amount, ok := units[unit]; ok {
		bill[item] += amount
		return true
	}
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if _, ok := bill[item]; ok {
		if amount, ok := units[unit]; ok {
			if bill[item] >= amount {
				bill[item] -= amount
				if bill[item] == 0 {
					delete(bill, item)
				}
				return true
			}
		}
		return false
	}
	return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemQuantity, ok := bill[item]
	return itemQuantity, ok
}
