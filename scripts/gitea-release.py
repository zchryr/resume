import json
import requests
import argparse
import datetime
import os

true = True
false = False

parser = argparse.ArgumentParser(description="Arguments.")
parser.add_argument('-release', required=True)
args = parser.parse_args()

repoOwner = os.environ.get('REPO_OWNER')
repoOwner = repoOwner.rstrip("\n")

repoName = os.environ.get('REPO_NAME')
repoName = repoName.rstrip("\n")

releaseAPIURL = "https://git.rohrbach.tech/api/v1/repos/" + repoOwner + "/" + repoName + "/releases"
apiAuth = "?access_token=" + os.environ.get('GITEA_API_TOKEN')

# Create a release.
def createRelease():
    today = datetime.datetime.now()
    date = "Resume-" + today.strftime("%b-%d-%Y")
    time = today.strftime("%m-%d-%Y_%H-%M-%S")

    reqBody = {
        "body": "",
        "draft": false,
        "name": date,
        "prerelease": false,
        "tag_name": time,
        "target_commitish": "master"
    }

    r = requests.post(releaseAPIURL + apiAuth, json=reqBody)

    if str(r.status_code) == "201":
        print("Release created")
        responseJson = json.loads(r.content.decode('utf-8'))
        return responseJson['id']
    else:
        print("Release creation failed.")
        exit(1)

if args.release == "true":
    createRelease()
else:
    print("No Gitea release was created..")