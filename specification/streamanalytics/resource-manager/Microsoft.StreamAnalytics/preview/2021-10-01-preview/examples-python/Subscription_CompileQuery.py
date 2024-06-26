from azure.identity import DefaultAzureCredential
from azure.mgmt.streamanalytics import StreamAnalyticsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-streamanalytics
# USAGE
    python subscription_compile_query.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = StreamAnalyticsManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="56b5e0a9-b645-407d-99b0-c64f86013e3d",
    )

    response = client.subscriptions.compile_query(
        location="West US",
        compile_query={
            "compatibilityLevel": "1.2",
            "functions": [
                {
                    "bindingType": "Microsoft.StreamAnalytics/JavascriptUdf",
                    "inputs": [{"dataType": "any", "isConfigurationParameter": None}],
                    "name": "function1",
                    "output": {"dataType": "bigint"},
                    "type": "Scalar",
                }
            ],
            "inputs": [{"name": "input1", "type": "Stream"}],
            "jobType": "Cloud",
            "query": "SELECT\r\n    *\r\nINTO\r\n    [output1]\r\nFROM\r\n    [input1]",
        },
    )
    print(response)


# x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Subscription_CompileQuery.json
if __name__ == "__main__":
    main()
