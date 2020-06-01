try:
    import pymongo
    from pymongo import MongoClient
    import pandas as pd
    import json
except Exception as e:
    print("Some Modules are Missing")
    
class MongoDB(object):

    def __init__(self, dBName=None, collectionName=None):
    
        self.dBName = dBName
        self.collectionName = collectionName
        
        self.client = MongoClient("localhost", 27017, maxPoolSize=50)
        
        self.DB = self.client[self.dBName]
        self.collection = self.DB[self.collectionName]
        
    def InsertData(self, path=None):
        """
        :param path: Path os csv File
        :return: None
        """
        
        df = pd.read_csv(path)
        data = df.to_dict('records')
        
        self.collection.insert_many(data, ordered=False)
        print("All the Data has been Exported to Mongo DB Server ...")

if __name__ == "__main__":
    mongodb = MongoDB(dBName = 'Dataset', collectionName = 'orders')
    mongodb.InsertData(path="task_data/Test task - Postgres - orders.csv")

    mongodb = MongoDB(dBName = 'Dataset', collectionName = 'order_items')
    mongodb.InsertData(path="task_data/Test task - Postgres - order_items.csv")

    mongodb = MongoDB(dBName = 'Dataset', collectionName = 'deliveries')
    mongodb.InsertData(path="task_data/Test task - Postgres - deliveries.csv")
    
    mongodb = MongoDB(dBName = 'Dataset', collectionName = 'customers')
    mongodb.InsertData(path="task_data/Test task - Mongo - customers.csv")
    
    mongodb = MongoDB(dBName = 'Dataset', collectionName = 'customer_companies')
    mongodb.InsertData(path="task_data/Test task - Mongo - customer_companies.csv")