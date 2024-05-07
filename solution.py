


def addition(l1,l2):
    carry = 0
    result = []
    for i in range(max(len(l1),len(l2))+1):
        tmp = (l1[i] if i < len(l1) else 0 )+ (l2[i] if i < len(l2) else 0 )+ carry
        carry,digit = divmod(tmp,10)
        result.append(digit)
    return result


l1_list = [9,9,9,9,9,9,9,1,2,3,4,5,6,7,8,9,0]
l2_list = [9,9,9,9]

l1_list.reverse()
l2_list.reverse()

result = addition(l1_list,l2_list)
print(result)