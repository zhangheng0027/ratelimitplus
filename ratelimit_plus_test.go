package ratelimit

import (
	gc "gopkg.in/check.v1"
)

func (rateLimitSuite) TestPlus(c *gc.C) {

	b1 := NewBucketWithRate(1, 100)
	b2 := NewBucketWithRate(1, 100)
	b3 := NewBucketWithRate(1, 100)

	b1.AddUpstream(b2)
	b1.AddUpstream(b3)

	b1.Take(2)

	c.Assert(b2.availableTokens, gc.Equals, int64(98))
	c.Assert(b3.availableTokens, gc.Equals, int64(98))
}

func (rateLimitSuite) TestPlus1(c *gc.C) {

	b1 := NewBucketWithRate(1, 100)
	b2 := NewBucketWithRate(1, 30)
	b3 := NewBucketWithRate(1, 100)

	b1.AddUpstream(b2)
	b1.AddUpstream(b3)
	b1.bucketPlus.SetControlModel(ParallelControl)
	b1.TakeMaxDuration(40, 0)

	c.Assert(b2.availableTokens, gc.Equals, int64(30))
	c.Assert(b3.availableTokens, gc.Equals, int64(60))
}

func (rateLimitSuite) TestPlus2(c *gc.C) {

	b1 := NewBucketWithRate(1, 100)
	b2 := NewBucketWithRate(1, 100)
	b3 := NewBucketWithRate(1, 100)

	b1.AddUpstream(b2)
	b2.AddUpstream(b3)

	b1.Take(2)
	b2.Take(2)
	b3.Take(2)

	c.Assert(b1.availableTokens, gc.Equals, int64(98))
	c.Assert(b2.availableTokens, gc.Equals, int64(96))
	c.Assert(b3.availableTokens, gc.Equals, int64(94))
}

func (rateLimitSuite) TestPlus3(c *gc.C) {

	b1 := NewBucketWithRate(1, 100)
	b2 := NewBucketWithRate(1, 100)
	b3 := NewBucketWithRate(1, 100)

	b1.AddUpstream(b3)
	b2.AddUpstream(b3)

	b1.Take(2)
	b2.Take(2)
	b3.Take(2)

	c.Assert(b1.availableTokens, gc.Equals, int64(98))
	c.Assert(b2.availableTokens, gc.Equals, int64(98))
	c.Assert(b3.availableTokens, gc.Equals, int64(94))
}