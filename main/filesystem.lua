
function test01() 
    file = io.open("test1_lua.txt", "w+")
    for i = 0, 10, 1 do
        file:write("Hello lua file io system\n")
    end
    file:close()
end


function test02() 
    file = io.open("test1_lua.txt", "a+")
    file:write("ABC!ENGLISH\n")
    file:close()
end

function test03()
    file = io.open("test1_lua.txt", "r")
    print(file:read("*a"))
    file:close()
    
    file = io.open("test1_lua.txt", "a+")

    for i = 0, 5, 1 do
        file:write("This line is written by lua\n")
    end
    file:close()
end

function test04()
    file = io.open("test1_lua.txt", "r")
    file1 = io.open("test1_lua1.txt", "w+")
    file1:write(file:read("*a"))
    file:close()
    file1:close()
end

function exists(a) 
    local file = io.open(a, "rb")
    if file then 
        file:close()
    end
    return file ~= nil
end

print(exists("hhh"))