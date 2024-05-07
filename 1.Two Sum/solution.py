

def two_number(number_list,target):
    for i,num in enumerate(number_list):
        missing = target - num
        if missing in number_list and missing != num:
            missing_index = number_list.index(missing)
            print(i,missing_index)

two_number([3,2,4],6)