


def addition(l1,l2)
    carry = 0
    result = []
    for i in (0...[l1.size,l2.size].max) do
        digit1 = 
        tmp = (i < l1.size ? l1[i] : 0) +(i < l2.size ? l2[i] : 0) + carry
        carry,digit = tmp.divmod(10)
        result.append(digit)
    end
    result.append(carry) if carry > 0
    return(result)
end

l1_list = [9,9,9,9,9,9,9]
l2_list = [9,9,9,9]

result = addition(l1_list.reverse(),l2_list.reverse())

puts "#{result}"