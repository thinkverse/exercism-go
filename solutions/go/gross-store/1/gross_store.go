package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := make(map[string]int)

	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728

	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if amount, ok := units[unit]; ok {
		bill[item] += amount
		return ok
	}

	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if quantity, ok := bill[item]; ok {
		if amount, ok := units[unit]; ok {
			if (quantity - amount) < 0 {
				return false
			}

			if (quantity - amount) == 0 {
				delete(bill, item)
			} else {
				bill[item] -= amount
			}

			return true
		}
	}

	return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if quantity, ok := bill[item]; ok {
		return quantity, ok
	}

	return 0, false
}
