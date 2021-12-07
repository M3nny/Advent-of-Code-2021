crabs = [16,1,2,0,4,2,7,1,2,14]
def median(array)
    return nil if array.empty?
    sorted = array.sort
    len = sorted.length
    (sorted[(len - 1) / 2] + sorted[len / 2]) / 2
    end
  
middle = median crabs
fuel = 0

crabs.each{
    |i|

    if i < middle
        fuel += (middle - i)
    elsif i > middle
        fuel += (i - middle)
    end
}

puts "the horizontal position that costs the least is: #{middle}"
puts "This costs a total of #{fuel} L of fuel"