package model

import "time"

type ContentKeyGenerator struct {
	PeriodList     *[]string
	KeysProperties *[]string
	KidMappings    *[]string
	Misc           string
}

type IContentId interface {
	ProcessPeriodList(startTime *time.Time, endTime *time.Time)
	GetKeyGenerator() KeyGenerator
	Enhance()
}

type KeyGenerator ContentKeyGenerator
type LinearGenerator KeyGenerator
type VodGenerator KeyGenerator

func NewKeyGenerator(startTime *time.Time, endTime *time.Time) IContentId {
	if startTime == nil && endTime == nil {
		tmp := VodGenerator{}
		tmp.Misc = "vod"
		return &tmp
	} else {
		tmp := LinearGenerator{}
		tmp.Misc = "linear"
		return &tmp
	}
}

func (c *VodGenerator) ProcessPeriodList(startTime *time.Time, endTime *time.Time) {
	var periodList []string

	period := "newPeriod"
	periodList = append(periodList, period)
	c.PeriodList = &periodList
	return
}

func (c *LinearGenerator) ProcessPeriodList(startTime *time.Time, endTime *time.Time) {

}

func (c *LinearGenerator) GetKeyGenerator() KeyGenerator {
	return KeyGenerator(*c)
}

func (c *VodGenerator) GetKeyGenerator() KeyGenerator {
	return KeyGenerator(*c)
}

func (c *LinearGenerator) Enhance() {
	return
}

func (c *VodGenerator) Enhance() {
	(*c.PeriodList)[0] = "yes"
	c.Misc = "no"
	return
}
