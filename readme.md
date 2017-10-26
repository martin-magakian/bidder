#build image
docker build -t martinmagakian/bidder .
docker run -it --add-host redis_db:192.168.99.100 -p 3001:3001 martinmagakian/bidder

#run tools
cd vendor/beeswax/tools/bid/bidding_agent_requests_generator/bid
./bidding_agent_requests_generator ../sample_bid_request_1.txt  http://192.168.99.100:3001/ --path-to-responses-file bid_response.txt --header-secret some-secret --log-level debug
