from azure.identity import DefaultAzureCredential
from azure.mgmt.streamanalytics import StreamAnalyticsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-streamanalytics
# USAGE
    python subscription_test_input.py

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

    response = client.subscriptions.begin_test_input(
        location="West US",
        test_input={
            "input": {
                "properties": {
                    "datasource": {
                        "properties": {
                            "container": "state",
                            "dateFormat": "yyyy/MM/dd",
                            "pathPattern": "{date}/{time}",
                            "sourcePartitionCount": 16,
                            "storageAccounts": [{"accountKey": "someAccountKey==", "accountName": "someAccountName"}],
                            "timeFormat": "HH",
                        },
                        "type": "Microsoft.Storage/Blob",
                    },
                    "serialization": {"properties": {"encoding": "UTF8", "fieldDelimiter": ","}, "type": "Csv"},
                    "type": "Stream",
                }
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Subscription_TestInput.json
if __name__ == "__main__":
    main()
