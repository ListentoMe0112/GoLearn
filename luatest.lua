-- function test (n1, n2)
--     return n1 + n2
-- end

-- function test01(test, n3)
--     return test(10, 20) + n3
-- end

-- print(test(1,2))

-- print(test01(test, 1))


-- function sum (n1, ...)
--     print(select('#', ...))
--     for k,v in ipairs({...}) do
--         n1 = n1 + v
--     end
--     return n1
-- end

-- print(sum(10, 201, 239, 31, 321, 43, 54, 31))

-- function acc()
--     local str = "hello"
--     local cur = cur or 10
--     return function (x)
--         cur = cur + x
--         str = str..cur
--         return cur, str
--     end
-- end

-- function makeSuffix(suffix)
--     local m_suffix = suffix
--     return function (fileName)
--         return fileName..m_suffix
--     end
-- end

-- a1 = acc()
-- print(a1(2))
-- print(a1(2))
-- print(a1(2))

-- jpg = makeSuffix(".jpg")
-- print(jpg("hello"))


-- Account = {}

-- function Account:withdraw (v)
-- 	self.balance = self.balance - v
-- end
 
-- function Account:deposit (v)
-- 	self.balance = self.balance + v
-- end

-- function Account:create(x)
--     local o = {balance = x or 0}
--     metatable = {}
--     metatable.__index = Account
--     metatable.__len = function() 
--         return 5
--     end
--     setmetatable(o, metatable)
--     return o
-- end



-- acc1 = Account:create()
-- print(#acc1)
-- function acc1:deposit (v)
-- 	self.balance = self.balance + v
--     print("balance: ", self.balance)
-- end
-- acc1:deposit(10)

-- mytable = {key1 = "value1"}
-- setmetatable(mytable, {
--     __newindex = function(mytable, key, value)
--         rawset(mytable, key, "\""..value.."\"")
--     end
-- })

-- mytable.key1 = "new value"
-- mytable.key2 = 4
-- print(mytable.key1,mytable.key2)

for k,v in pairs(arg) do
    print(k, v)
end