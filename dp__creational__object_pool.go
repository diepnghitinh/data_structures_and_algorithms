package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
)

/*

object pool là mẫu thiết kế trong đó một nhóm các đối tượng được khởi tạo và tạo trước đó và được giữ trong một nhóm.
Khi cần thiết, Client có thể yêu cầu một đối tượng từ nhóm,
sử dụng nó và trả lại cho nhóm. Đối tượng trong Pool không bao giờ bị drop.

*/

type iPoolObject interface {
	getID() string //This is any id which can be used to compare two different pool objects
}

type pool struct {
	idle     []iPoolObject
	active   []iPoolObject
	capacity int
	mulock   *sync.Mutex
}

//InitPool Initialize the pool
func initPool(poolObjects []iPoolObject) (*pool, error) {
	if len(poolObjects) == 0 {
		return nil, fmt.Errorf("Cannot craete a pool of 0 length")
	}
	active := make([]iPoolObject, 0)
	pool := &pool{
		idle:     poolObjects,
		active:   active,
		capacity: len(poolObjects),
		mulock:   new(sync.Mutex),
	}
	return pool, nil
}

func (p *pool) loan() (iPoolObject, error) {
	p.mulock.Lock()
	defer p.mulock.Unlock()
	if len(p.idle) == 0 {
		return nil, fmt.Errorf("No pool object free. Please request after sometime")
	}
	obj := p.idle[0]
	p.idle = p.idle[1:]
	p.active = append(p.active, obj)
	fmt.Printf("Loan Pool Object with ID: %s\n", obj.getID())
	return obj, nil
}

func (p *pool) receive(target iPoolObject) error {
	p.mulock.Lock()
	defer p.mulock.Unlock()
	err := p.remove(target)
	if err != nil {
		return err
	}
	p.idle = append(p.idle, target)
	fmt.Printf("Return Pool Object with ID: %s\n", target.getID())
	return nil
}

func (p *pool) remove(target iPoolObject) error {
	currentActiveLength := len(p.active)
	for i, obj := range p.active {
		if obj.getID() == target.getID() {
			p.active[currentActiveLength-1], p.active[i] = p.active[i], p.active[currentActiveLength-1]
			p.active = p.active[:currentActiveLength-1]
			return nil
		}
	}
	return fmt.Errorf("Targe pool object doesn't belong to the pool")
}

type connection struct {
	id string
}

func (c *connection) getID() string {
	return c.id
}

func dp__creational__object_pool() {

	connections := make([]iPoolObject, 0)
	for i := 0; i < 3; i++ {
		c := &connection{id: strconv.Itoa(i)}
		connections = append(connections, c)
	}
	pool, err := initPool(connections)
	if err != nil {
		log.Fatalf("Init Pool Error: %s", err)
	}
	conn1, err := pool.loan()
	if err != nil {
		log.Fatalf("Pool Loan Error: %s", err)
	}
	conn2, err := pool.loan()
	if err != nil {
		log.Fatalf("Pool Loan Error: %s", err)
	}
	pool.receive(conn1)
	pool.receive(conn2)

}
