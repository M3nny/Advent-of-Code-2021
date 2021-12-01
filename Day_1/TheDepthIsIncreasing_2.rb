#!/usr/bin/ruby
measurements = Array.new
x, increase, cont = 0, 0, 1
prev = 0
file = File.open("input.txt", 'r')

while !file.eof?
    measurements[x] = file.readline.to_i
    x += 1 
end

measurements[0..-1].each_slice(3) {
        |n|

        if n.length == 3
            puts "#{n.sum} (N/A - no previous sum)"
            prev = n.sum
        end
        break
    }

while cont < x
    measurements[cont..-1].each_slice(3) {
        |n|

        if n.length == 3
            if n.sum > prev
                puts "#{n.sum} (increased)"
                increase += 1
            elsif n.sum == prev
                puts "#{n.sum} (no change)"
            
            else
                puts "#{n.sum} (decreased)"
            end

            prev = n.sum
        end
        break
    }
    cont += 1
end

puts "there are #{increase} sums that are larger than the previous sum."