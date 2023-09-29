import requests
import json
from google.rpc.status_pb2 import Status

class MailExecutor(Executor):
    def run_query(self, query):
        q = json.loads(query.input)
        headers = {'Accept': 'application/json'}
        r = requests.get(f'https://api.mailcheck.ai/email/{q["user.email"]}', headers=headers)
        results = r.json()

        response = QueryResponse(status=Status(code=0, message="Success"))
        response.results = json.dumps({"disposable": results["disposable"], "alias": results["alias"]}) 
        return response

{'executor': MailExecutor, 'types': MailExecutor.supported_types}
