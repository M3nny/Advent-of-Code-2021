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
            #puts "#{b} (increased)" debug
            increase += 1
        else
            #puts "#{b} (decreased)" debug
        end
    end
}

puts "there are  #{increase} measurements that are larger than the previous measurement."