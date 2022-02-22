package user

import (
	"errors"
	"reflect"
	"strconv"
)

type property struct {
	Name  string
	value int
	min   int
	max   int
}

func (ppt *property) set(value int) error {
	ppt.value = value
	if ppt.value < ppt.min {
		ppt.value = ppt.min
		return errors.New("less than min")
	} else if ppt.value > ppt.max {
		ppt.value = ppt.max
		return errors.New("more than max")
	}
	return nil
}

func (ppt *property) change(value int) error {
	temp := ppt.value + value
	err := ppt.set(temp)
	return err
}

func (ppt *property) BiggerThan(value int) bool {
	if ppt.value > value {
		return true
	} else {
		return false
	}
}

func (ppt *property) Equal(value int) bool {
	if ppt.value == value {
		return true
	} else {
		return false
	}
}

type intProperty struct {
	property
}

func (ppt *intProperty) Value() int {
	return ppt.value
}

//----------------------------------

type switchProperty struct {
	property
	vnMap map[int]string
	nvMap map[string]int
}

func (ppt *switchProperty) set(option string) error {
	if _, ok := ppt.nvMap[option]; ok {
		ppt.value = ppt.nvMap[option]
		return nil
	} else {
		return errors.New(option + " not exist")
	}
}

func (ppt *switchProperty) BiggerThan(option string) bool {
	if _, ok := ppt.nvMap[option]; ok {
		if ppt.value > ppt.nvMap[option] {
			return true
		} else {
			return false
		}
	} else {
		return true //大于不存在的选项
	}
}

func (ppt *switchProperty) Equal(option string) bool {
	if _, ok := ppt.nvMap[option]; ok {
		if ppt.value == ppt.nvMap[option] {
			return true
		} else {
			return false
		}
	} else {
		return true //大于不存在的选项
	}
}

func (ppt *switchProperty) Value() string {
	return ppt.vnMap[ppt.value]
}

func NewIntProperty(name string, init, min, max int) intProperty {
	ppt := intProperty{}
	ppt.Name = name
	if max < min {
		min, max = max, min
	}
	ppt.min = min
	ppt.max = max
	ppt.set(init)
	return ppt
}

func NewSwitchProperty(name, init string, options []string) switchProperty {
	ppt := switchProperty{
		vnMap: map[int]string{},
		nvMap: map[string]int{},
	}
	ppt.Name = name
	offset := 0
	for v, option := range options {
		if _, ok := ppt.nvMap[option]; !ok {
			ppt.nvMap[option] = v - offset
		} else {
			offset += 1
		}
	}
	//无相同键
	for option, v := range ppt.nvMap {
		ppt.vnMap[v] = option
	}
	ppt.min = 0
	ppt.max = len(ppt.vnMap) - 1
	ppt.value = ppt.min
	ppt.set(init)
	return ppt
}

type User struct {
	intProperties    map[string]intProperty
	switchProperties map[string]switchProperty
	permission       map[string]string
}

func (user *User) AddSwitch(ppt switchProperty, perm string) {
	if _, ok := user.permission[ppt.Name]; !ok {
		user.switchProperties[ppt.Name] = ppt
		user.permission[ppt.Name] = perm
	}
}

func (user *User) AddInt(ppt intProperty, perm string) {
	if _, ok := user.permission[ppt.Name]; !ok {
		user.intProperties[ppt.Name] = ppt
		user.permission[ppt.Name] = perm
	}
}

func (user *User) SetSwitch(name, option string, permUser User) error {
	if _, ok := user.switchProperties[name]; ok {
		permppt := permUser.switchProperties["__level__"]
		if reflect.DeepEqual(permppt.vnMap, user.switchProperties["__level__"].vnMap) {
			if permppt.BiggerThan(user.permission[name]) || permppt.Equal(user.permission[name]) { //许可用户权限足够
				userppt := user.switchProperties[name]
				err := userppt.set(option)
				user.switchProperties[name] = userppt
				return err
			} else {
				return errors.New("permission denied")
			}
		} else {
			return errors.New("invalid permission user")
		}
	} else {
		return errors.New(name + " not exist")
	}
}

func (user *User) ChangeInt(name string, change int, permUser User) error {
	userppt := user.intProperties[name]
	err := user.SetInt(name, userppt.Value()+change, permUser)
	return err
}

func (user *User) SetInt(name string, change int, permUser User) error {
	if _, ok := user.intProperties[name]; ok {
		permppt := permUser.switchProperties["__level__"]
		if reflect.DeepEqual(permppt.vnMap, user.switchProperties["__level__"].vnMap) {
			if permppt.BiggerThan(user.permission[name]) || permppt.Equal(user.permission[name]) { //许可用户权限足够
				userppt := user.intProperties[name]
				err := userppt.set(change)
				user.intProperties[name] = userppt
				return err
			} else {
				return errors.New("permission denied")
			}
		} else {
			return errors.New("invalid permission user")
		}
	} else {
		return errors.New(name + " not exist")
	}
}

func (user *User) GetIntProperty(name string) (intProperty, error) {
	if _, ok := user.intProperties[name]; ok {
		return user.intProperties[name], nil
	} else {
		return intProperty{}, errors.New(name + " not exist")
	}
}

func (user *User) GetSwitchProperty(name string) (switchProperty, error) {
	if _, ok := user.switchProperties[name]; ok {
		return user.switchProperties[name], nil
	} else {
		return switchProperty{}, errors.New(name + " not exist")
	}
}

func (user *User) Show() {
	println("\nInt Properties:")
	for pptName, ppt := range user.intProperties {
		println(pptName + ": " + strconv.Itoa(ppt.Value()) + " _/" + user.permission[pptName])
	}
	println("\nSwitch Properties:")
	for pptName, ppt := range user.switchProperties {
		println(pptName + ": " + ppt.Value() + " _/" + user.permission[pptName])
	}
}

func NewUser(level switchProperty, perm string) User { //修改此用户等级所需权限
	user := User{
		intProperties:    map[string]intProperty{},
		switchProperties: map[string]switchProperty{},
		permission:       map[string]string{},
	}
	level.Name = "__level__"
	user.AddSwitch(level, perm)
	return user
}
