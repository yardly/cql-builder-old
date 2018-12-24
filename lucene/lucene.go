package lucene

import (
	"encoding/json"

	"github.com/gocql/gocql"
)

type Option func(*Options)

type Options struct {
	Column string
	Lucene *Lucene
}

type Rule interface {
	Lucene() bool
}

type Lucene struct {
	Filters        []Rule `json:"filter,omitempty"`
	Queries        []Rule `json:"query,omitempty"`
	Sorts          []Rule `json:"sort,omitempty"`
	PerformRefresh bool   `json:"refresh,omitempty"`
}

func (v *Lucene) MarshalCQL(_ gocql.TypeInfo) ([]byte, error) {
	return json.Marshal(v)
}

func NewOptions(os ...Option) Options {
	opts := Options{
		Column: "lucene",
		Lucene: new(Lucene),
	}
	for _, o := range os {
		o(&opts)
	}
	return opts
}

func Column(s string) Option {
	return func(o *Options) {
		o.Column = s
	}
}

func Filter(r ...interface{}) Option {
	return func(o *Options) {
		for _, e := range r {
			if i, ok := e.(Rule); ok && i != nil {
				o.Lucene.Filters = append(o.Lucene.Filters, i)
			}
		}
	}
}

func Query(r ...interface{}) Option {
	return func(o *Options) {
		for _, e := range r {
			if i, ok := e.(Rule); ok && i != nil {
				o.Lucene.Queries = append(o.Lucene.Queries, i)
			}
		}
	}
}

func Sort(r ...interface{}) Option {
	return func(o *Options) {
		for _, e := range r {
			if i, ok := e.(Rule); ok && i != nil {
				o.Lucene.Sorts = append(o.Lucene.Sorts, i)
			}
		}
	}
}

func Refresh() Option {
	return func(o *Options) {
		o.Lucene.PerformRefresh = true
	}
}
