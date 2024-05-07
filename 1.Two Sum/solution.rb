

def two_number(number_list,target)
    number_list.each_with_index do |num,i|
        missing = target - num
        missing_index = number_list.index(missing)
        if missing_index != i
            puts "#{i} #{missing_index}"
        end
    end
    
end


two_number([3,2,4],6)