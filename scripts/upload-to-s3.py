import logging
import boto3
import argparse
from botocore.exceptions import ClientError

def uploadFile(fileName, bucket, objectName=None):
    # If S3 objectName was not specified, use fileName
    if objectName is None:
        objectName = fileName

    # Upload the file
    s3_client = boto3.client('s3')
    try:
        response = s3_client.upload_file(fileName, bucket, objectName)
    except ClientError as e:
        logging.error(e)
        print("S3 Upload Failed.")
        return False
    print("S3 Upload Successful.")
    return True

uploadFile('Zachary-Rohrbach-Resume.pdf', 'rohrbach')