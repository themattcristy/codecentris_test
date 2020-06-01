Requirements:
python
golang (will be using localhost:12345)
nodejs
mongodb (expecting to run at localhost:27017)



Steps:

1.Export the csv files to Mongodb
	- open and connect the mongodb
	- ensure that you have pymongo and pandas python packages
	- if not, open cmd with the project folder directory and type "pip install pymongo" and
	  pip install pandas.
	- after installing the packages type again in cmd "python export_csv_to_db.py"
	- In you mongodb you should see a new database names Dataset containing 5 collection.

2.Run the backend server
	 -ensure that you have following golang dependencies: github.com/gorilla/mux & go.mongodb.org/mongo-driver/mongo
	 -you can obtain that by typing "go get github.com/gorilla/mux" and "go get go.mongodb.org/mongo-driver/mongo" in cmd
	 -In cmd, type go build "go run db_api.go" and you should able to see "Server Running..." means
	  the backend is successfully running.

3.Run the app
	-open another cmd and cd to orders_app directory
	-type in cmd "npm install" to install the required packages
	-after that run the project by typing "npm run serve"
	-there should be a response that app is running at local with the following url. 
	-Copy the Url and paste it in the browser and there you should be able to see the application.