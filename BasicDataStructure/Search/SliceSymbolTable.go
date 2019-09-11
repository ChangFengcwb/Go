package Search

type element struct {
	key   interface{}
	value interface{}
}

func newElement(key interface{}, value interface{}) *element {
	return &element{key: key, value: value}
}

func (this *element) GetKey() interface{} {
	return this.key
}
func (this *element) GetValue() interface{} {
	return this.value
}
func (this *element) SetKey(key interface{}) {
	this.key = key
}
func (this *element) SetValue(value interface{}) {
	this.value = value
}

type SliceSymbolTable struct {
	elements []*element
	size     int
}

func NewSliceSymbolTable() *SliceSymbolTable {
	return &SliceSymbolTable{
		elements: make([]*element, 0),
		size:     0,
	}
}

func (this *SliceSymbolTable) rank(key interface{}) int {
	if this.size <= 0 || key == nil {
		return -1
	}
	left, right := 0, this.size
	for left < right {
		mid := left + ((right - left) >> 1)
		switch key.(type) {
		case string:
			if this.elements[mid].GetKey().(string) < key.(string) {
				left = mid + 1
			} else if this.elements[mid].GetKey().(string) > key.(string) {
				right = mid - 1
			} else {
				return mid
			}
		case int:
			if this.elements[mid].GetKey().(int) < key.(int) {
				left = mid + 1
			} else if this.elements[mid].GetKey().(int) > key.(int) {
				right = mid - 1
			} else {
				return mid
			}
		case float64:
			if this.elements[mid].GetKey().(float64) < key.(float64) {
				left = mid + 1
			} else if this.elements[mid].GetKey().(float64) > key.(float64) {
				right = mid - 1
			} else {
				return mid
			}
		}
	}
	return -1
}
func (this *SliceSymbolTable) Put(key, value interface{}) {
	if this.Contains(key) {
		this.elements[this.rank(key)].SetValue(value)
	} else {
		if this.size == 0 {
			this.elements = append(this.elements, newElement(key, value))
			this.size++
		} else {
			for i := 0; i < this.size; i++ {
				switch key.(type) {
				case string:
					if this.elements[i].GetKey().(string) < key.(string) {
						this.elements = append(this.elements, nil)
						copy(this.elements[i+1:], this.elements[i:])
						this.elements[i] = newElement(key, value)
						this.size++
						return
					}
				case int:
					if this.elements[i].GetKey().(int) < key.(int) {
						this.elements = append(this.elements, nil)
						copy(this.elements[i+1:], this.elements[i:])
						this.elements[i] = newElement(key, value)
						this.size++
						return
					}
				case float64:
					if this.elements[i].GetKey().(float64) < key.(float64) {
						this.elements = append(this.elements, nil)
						copy(this.elements[i+1:], this.elements[i:])
						this.elements[i] = newElement(key, value)
						this.size++
						return
					}
				}
			}
			this.elements = append(this.elements, newElement(key, value))
			this.size++
		}
	}
}

func (this *SliceSymbolTable) Get(key interface{}) interface{} {
	index := this.rank(key)
	if index != -1 {
		return this.elements[index].GetValue()
	}
	return nil
}

func (this *SliceSymbolTable) Delete(key interface{}) {
	index := this.rank(key)
	if index != -1 {
		this.elements = append(this.elements[:index], this.elements[index+1:]...)
		this.size--
	}
}

func (this *SliceSymbolTable) Contains(key interface{}) bool {
	return this.rank(key) != -1
}

func (this *SliceSymbolTable) IsEmpty() bool {
	return this.size == 0
}

func (this *SliceSymbolTable) Size() int {
	return this.size
}

func (this *SliceSymbolTable) KeySet() []interface{} {
	keys := make([]interface{}, this.size)
	for i := 0; i < this.size; i++ {
		keys[i] = this.elements[i].GetKey()
	}
	return keys
}
