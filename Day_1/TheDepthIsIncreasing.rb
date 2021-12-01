#!/usr/bin/ruby
measurements = Array.new
x, increase=0, 0
file = File.open("input.txt", 'r')
while !file.eof?
    measurements[x] = file.readline.to_i
    x += 1 
end

measurements.each_cons(2) { 
    |a, b|
    if b != 0
        if b > a
            puts "#{b} (increased)" 
            increase += 1
        elsif b == a 
            puts "#{b} (no change)" 
        else
            puts "#{b} (decreased)" 
        end
    end
}

puts "there are  #{increase} measurements that are larger than the previous measurement."