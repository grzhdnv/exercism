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
	if i, ok := units[unit]; ok {
		bill[item] += i
		return true
	} else {
		return false
	}
}

// RemoveItem removes an item from customer bill.
// Return false if the given item is not in the bill
// Return false if the given unit is not in the units map.
// Return false if the new quantity would be less than 0.
// If the new quantity is 0, completely remove the item from the bill then return true.
// Otherwise, reduce the quantity of the item and return true.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	unitQty, ok := units[unit]
	if !ok {
		return false
	}

	itemQty, ok := bill[item]
	if !ok {
		return false
	}

	newQty := itemQty - unitQty
	if newQty < 0 {
		return false
	}

	if newQty == 0 {
		delete(bill, item)
		return true
	}

	bill[item] = newQty
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {

	quantity, in := bill[item]

	return quantity, in
}
