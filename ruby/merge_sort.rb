def generate_unsorted(size, upper)
  return size.times.map{Random.rand(upper)} 
end

def merge_sort(unsorted) 
  if (unsorted.length <= 1)
    return unsorted
  end
  pivot = unsorted.length / 2
  sub_arr_1 = merge_sort(unsorted[0, pivot])
  sub_arr_2 = merge_sort(unsorted[pivot, unsorted.length-1])
  sub_ptr_1 = 0
  sub_ptr_2 = 0
  for i in 0..unsorted.length-1 
    if (sub_ptr_1 >= sub_arr_1.length )
	  unsorted[i] = sub_arr_2[sub_ptr_2]
	  sub_ptr_2 += 1
	elsif (sub_ptr_2 >= sub_arr_2.length)
	  unsorted[i] = sub_arr_1[sub_ptr_1]
	  sub_ptr_1 += 1
    elsif (sub_arr_1[sub_ptr_1] <= sub_arr_2[sub_ptr_2])
	  unsorted[i] = sub_arr_1[sub_ptr_1]
	  sub_ptr_1 += 1
	else 
      unsorted[i] = sub_arr_2[sub_ptr_2]
	  sub_ptr_2 += 1
    end
  end
  return unsorted
end 

unsorted = generate_unsorted(10, 10)
puts 'Unsorted:'
puts unsorted
puts 'Sorted:'
puts merge_sort(unsorted)