package container

import (
	"errors"
	"log"
	"sync"

	"go.uber.org/dig"
)

const (
	// DefaultContainerName 默认容器名称
	DefaultContainerName string = "_default_"
)

var (
	instance *di
	once     sync.Once
)

// Container 容器
type Container struct {
	container *dig.Container
}

type di struct {
	mu         sync.RWMutex
	containers map[string]*Container // 容器列表
}

// Option 容器选项
type Option struct {
	Name string // 容器名字
}

// ProviderOption a provideOption modifies the default behavior of Provide.
type ProvideOption interface {
	// applyProvideOption(*provideOptions)
	// applyProvideOption()
}
type InvokeOption interface {
	// unimplemented()
}

// singleton 内部单例
func singleton() *di {
	once.Do(func() {
		instance = &di{containers: make(map[string]*Container)}
		instance.add(DefaultContainerName)
	})

	return instance
}

// New创建一个容器
func New(option ...Option) *Container {
	if len(option) > 0 {
		name := option[0].Name
		if name != "" {
			singleton().add(name)
			return Get(name)
		}
	}
	return Default()
}

//  获取默认容器
func Default() *Container {
	return Get(DefaultContainerName)
}

// Get 获取容器
func Get(name string) *Container {
	c := singleton()
	c.mu.RLock()
	val, ok := c.containers[name]
	c.mu.RUnlock()
	if ok {
		return val
	}
	panic(errors.New("not found named [" + name + "] container"))
}

// Delete 删除容器
func Delete(name string) {
	c := singleton()
	c.remove(name)
}

func MustProvide(ctor interface{}, opts ...ProvideOption) {
	if err := Default().Provide(ctor); err != nil {
		panic(err)
	}
}
func Provide(ctor interface{}, opts ...ProvideOption) error {
	return Default().Provide(ctor)
}

func MustInvoke(fn interface{}, opts ...InvokeOption) {
	if err := Default().Invoke(fn); err != nil {
		panic(err)
	}
}
func Invoke(fn interface{}, opts ...InvokeOption) error {
	return Default().Invoke(fn)
}

// DI
//----------------------------------------------------------------------------------
func (d *di) add(name string) {
	dc := dig.New()
	c := &Container{container: dc}
	d.mu.Lock()
	d.containers[name] = c
	d.mu.Unlock()
}
func (d *di) remove(name string) {
	d.mu.Lock()
	delete(d.containers, name)
	d.mu.Unlock()
}

// Container
//----------------------------------------------------------------------------------
func (c *Container) Provide(ctor interface{}, opts ...ProvideOption) error {
	err := c.container.Provide(ctor)
	if err != nil {
		log.Printf("[container.Provide] err=%v\n", err)
	}
	return err
}

func (c *Container) Invoke(fn interface{}, opts ...InvokeOption) error {
	err := c.container.Invoke(fn)
	if err != nil {
		log.Printf("[container.Invoke] err=%v\n", err)
	}
	return err
}
