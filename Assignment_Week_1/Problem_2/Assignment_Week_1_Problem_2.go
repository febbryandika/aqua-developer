package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type Desc struct {
	Description string `json:"description"`
}

type Job struct {
	Title      string `json:"title"`
	WorkFrom   string `json:"work_from"`
	Department string `json:"department"`
}

type Cache struct {
	Aggr
	IsFilled bool
}

type Aggr struct {
	Desc Desc
	Jobs []Job
}

func caculateTime(start time.Time) {
	fmt.Println("dari calculate ", start)
	fmt.Printf("took %v\n", time.Since(start))
}

func fetchAPI(url string, obj interface{}) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, obj); err != nil {
		return err
	}

	return err
}

func (c *Cache) Aggregate() (*Cache, error) {
	defer caculateTime(time.Now())

	var wg sync.WaitGroup
	wg.Add(2)

	if !c.IsFilled {
		var desc Desc
		var descError error
		go func() {
			descError = fetchAPI("https://workspace-rho.vercel.app/api/description", &desc)

			wg.Done()
		}()

		var jobs []Job
		var jobsError error
		go func() {
			jobsError = fetchAPI("https://workspace-rho.vercel.app/api/jobs", &jobs)

			wg.Done()
		}()

		wg.Wait()

		if descError != nil {
			return &Cache{}, descError
		}
		if jobsError != nil {
			return &Cache{}, jobsError
		}

		c.Aggr.Desc = desc
		c.Aggr.Jobs = jobs
		c.IsFilled = true

		return c, nil
	} else {
		c.IsFilled = false
		return c, nil
	}
}

func main() {
	var c Cache
	_, err := c.Aggregate()
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%+v\n", c)
	_, err = c.Aggregate()
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%+v\n", c)
	_, err = c.Aggregate()
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%+v\n", c)
	_, err = c.Aggregate()
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%+v\n", c)
}
