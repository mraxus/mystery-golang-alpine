# encoding: utf-8

# https://stackoverflow.com/questions/1489183/colorized-ruby-output/16363159#16363159
class String
def green;          "\e[32m#{self}\e[0m" end
def magenta;        "\e[35m#{self}\e[0m" end
end

IMAGE_BUILD = "mamarcus.org/project:build"
IMAGE_ALPINE = "mamarcus.org/project:alpine"

desc 'Build the deployable image'
task :build_build do
  sh "docker build -f Dockerfile.build -t #{IMAGE_BUILD} ."
end
task :build_alpine do
  sh "docker build -f Dockerfile.alpine -t #{IMAGE_ALPINE} ."
end

desc 'Run docker-compose'
task :run do
    sh "docker-compose up -d postgres"
    print "Checking connection to postgres ('psql' must be installed)"
    psql_url = "postgres://postgres@#{`docker-compose port postgres 5432`.strip}?sslmode=disable"
    until system("psql -c 'SELECT 1' #{psql_url} > /dev/null 2>&1") do
        print "."
        sleep(1)
    end
    print "\nConnected\n\n"
    puts "Running service on build image. This should work fine:".green
    sh "docker-compose up build"

    print "\nBut: Running the alpine image; this will not connect nor quit/timeout for 10 min + ...\n".magenta
    sh "docker-compose up alpine"
end

desc 'Clean up docker-compose images'
task :clean do
  sh "docker-compose down"
end

task :default => [
	:build_build,
	:build_alpine,
	:run
]
