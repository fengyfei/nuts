/*
 * MIT License
 *
 * Copyright (c) 2017 SmartestEE Co., Ltd..
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2017/09/17        Yang Chenglong
 */

package test

import (
	"fmt"
	"testing"

	"github.com/fengyfei/nuts/kv"
	_ "github.com/fengyfei/nuts/kv/store/bbolt"
)

func TestAlias_Init(t *testing.T) {
	var alice *kv.Alias = new(kv.Alias)
	alice.Init("user.db")

	type S struct {
		A string
	}
	test := S{"test"}

	alice.Store.Put("user", 1234, "1234")
	v, _ := alice.Store.Get("user", 1234)
	fmt.Println(string(v))

	alice.Store.Put("user", "1234", 1234)
	v, _ = alice.Store.Get("user", "1234")
	fmt.Println(string(v))

	alice.Store.Put("user", test, "test")
	v, _ = alice.Store.Get("user", test)
	fmt.Println(string(v))

	alice.Store.Delete("user", test)
	v, _ = alice.Store.Get("user", test)
	fmt.Println(string(v))
}
