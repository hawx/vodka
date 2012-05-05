desc 'Minify js into a single file'
task :minify do
  require 'jsmin'

  files = %w(jquery.appear.min.js rainbow.min.js rainbow.vodka.js doc.js)

  out = 'rsc/min.js'
  res = ''

  files.each do |path|
    path = "rsc/#{path}"
    res << File.read(path) << "\n"
  end

  res = JSMin.minify(res)
  File.open(out, 'w') {|f| f.write(res) }
end
