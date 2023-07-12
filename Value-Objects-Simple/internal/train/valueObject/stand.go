package train

// 旅程
type Stand struct {
	Start string
	End   string
}

func CreatStand(start string, end string) Stand {
	return Stand{
		Start: start,
		End:   end,
	}
}
