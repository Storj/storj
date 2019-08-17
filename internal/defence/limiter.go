// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package defence

import (
	"context"
	"sync"
	"time"

	"golang.org/x/time/rate"

	"storj.io/storj/internal/sync2"
)

// Attacker stores information about attacker entity
type Attacker struct {
	limiter *rate.Limiter
	Expire  time.Time
	// TODO: could be interface{}
	key string

	IsBanned bool
}

func (attacker *Attacker) attack(attemptTime time.Time, banDuration time.Duration) bool {
	// is not cleared by Clear method of Limiter
	if attacker.IsBanned && attemptTime.After(attacker.Expire) {
		attacker.IsBanned = false
	}

	if attacker.IsBanned {
		return true
	}

	isBanned := !attacker.limiter.AllowN(attemptTime, 1)

	if isBanned {
		attacker.IsBanned = true
		attacker.Expire = attemptTime.Add(banDuration)
	}

	return isBanned
}

// Limiter is used to store and manage list of banned entities
type Limiter struct {
	attackers map[string]*Attacker

	// Attempts defines how many times attacker could perform an operation
	Attempts int
	// AttemptsPeriod defines period in which attempts will count. For example, 5 attempts per minute.
	AttemptsPeriod time.Duration
	BanDuration    time.Duration

	mu   sync.Mutex
	loop *sync2.Cycle
}

// NewLimiter is a constructor for Limiter
func NewLimiter(attempts int, attemptsPeriod, banDuration, clearPeriod time.Duration) *Limiter {
	return &Limiter{
		attackers:      map[string]*Attacker{},
		Attempts:       attempts,
		AttemptsPeriod: attemptsPeriod,
		BanDuration:    banDuration,
		loop:           sync2.NewCycle(clearPeriod),
	}
}

// Attack is use to add new fail attack
func (limiter *Limiter) Attack(key string) bool {
	limiter.mu.Lock()

	defer limiter.mu.Unlock()
	now := time.Now()

	// Try to retrieve the
	if _, found := limiter.attackers[key]; !found {
		limiter.attackers[key] = &Attacker{
			key:     key,
			limiter: rate.NewLimiter(rate.Every(limiter.BanDuration), limiter.Attempts),
			Expire:  now.Add(limiter.BanDuration),
		}
	}

	limit := limiter.attackers[key]

	return limit.attack(now, limiter.BanDuration)
}

// Banned returns the list of banned attackers
func (limiter *Limiter) Banned() []*Attacker {
	limiter.mu.Lock()

	defer limiter.mu.Unlock()

	var attackers []*Attacker

	for _, attacker := range limiter.attackers {
		if attacker.IsBanned {
			attackers = append(attackers, attacker)
		}
	}

	return attackers
}

// Find can be used to find an attacker by specified key
func (limiter *Limiter) Find(key string) (Attacker, bool) {
	limiter.mu.Lock()

	defer limiter.mu.Unlock()

	attacker, ok := limiter.attackers[key]
	return *attacker, ok
}

// CleanUp is used to clean all attackers whose ban is expired
func (limiter *Limiter) CleanUp(ctx context.Context) error {
	return limiter.loop.Run(ctx, func(ctx context.Context) error {
		limiter.cleanUpCallback()
		return nil
	})
}

func (limiter *Limiter) cleanUpCallback() {
	limiter.mu.Lock()

	defer limiter.mu.Unlock()

	for key, limit := range limiter.attackers {
		if time.Now().After(limit.Expire) {
			delete(limiter.attackers, key)
		}
	}
}

// Close should be used when limiter is no longer needed
func (limiter *Limiter) Close() {
	limiter.mu.Lock()

	defer limiter.mu.Unlock()

	limiter.loop.Stop()
}
