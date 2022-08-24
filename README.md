# zflag
format gin response && validator transfer

#### Usage
```
// We can omit c.Abort()
func Test(c *gin.Context) {
	h := new(dto.ParamTest)
	if err := c.ShouldBindJSON(&h); err != nil {
		zflag.Done(c, zflag.ParamErr(err))
	} else {
		if err := service.Test(c, h); err != nil {
			zflag.Done(c, zflag.ResponseErr(zflag.ErrQuery.WithMessage(err.Error())))
		} else {
			zflag.Done(c, zflag.Success(zflag.EmptyData))
		}
	}
}

// {"flag":1,"data":{},"message":"","detail":"","timestamp":""}
```

#### Transfer
```
func main() {
    s := gin.New()
    
    // transfer en-US -> zh-CN
    zflag.Transfer()
    
    s.GET("/test", Test)
    s.Run(":8080")
}
```