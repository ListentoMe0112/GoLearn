function test (n1, n2)
    return n1 + n2
end

function test01(test, n3)
    return test(10, 20) + n3
end

print(test(1,2))

print(test01(test, 1))


function sum (n1, ...)
    print(select('#', ...))
    for k,v in ipairs({...}) do
        n1 = n1 + v
    end
    return n1
end

print(sum(10, 201, 239, 31, 321, 43, 54, 31))

function acc()
    local str = "hello"
    local cur = cur or 10
    return function (x)
        cur = cur + x
        str = str..cur
        return cur, str
    end
end

function makeSuffix(suffix)
    local m_suffix = suffix
    return function (fileName)
        return fileName..m_suffix
    end
end

a1 = acc()
print(a1(2))
print(a1(2))
print(a1(2))

jpg = makeSuffix(".jpg")
print(jpg("hello"))