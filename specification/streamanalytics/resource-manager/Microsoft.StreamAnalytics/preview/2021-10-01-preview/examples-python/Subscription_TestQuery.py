from azure.identity import DefaultAzureCredential
from azure.mgmt.streamanalytics import StreamAnalyticsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-streamanalytics
# USAGE
    python subscription_test_query.py

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

    response = client.subscriptions.begin_test_query(
        location="West US",
        test_query={
            "diagnostics": {"path": "/pathto/subdirectory", "writeUri": "http://myoutput.com"},
            "streamingJob": {
                "location": "West US",
                "properties": {
                    "compatibilityLevel": "1.0",
                    "dataLocale": "en-US",
                    "eventsLateArrivalMaxDelayInSeconds": 5,
                    "eventsOutOfOrderMaxDelayInSeconds": 0,
                    "eventsOutOfOrderPolicy": "Drop",
                    "functions": [],
                    "inputs": [
                        {
                            "name": "inputtest",
                            "properties": {
                                "datasource": {"properties": {"payloadUri": "http://myinput.com"}, "type": "Raw"},
                                "serialization": {"properties": {"encoding": "UTF8"}, "type": "Json"},
                                "type": "Stream",
                            },
                        }
                    ],
                    "outputErrorPolicy": "Drop",
                    "outputs": [
                        {
                            "name": "outputtest",
                            "properties": {
                                "datasource": {"properties": {"payloadUri": "http://myoutput.com"}, "type": "Raw"},
                                "serialization": {"type": "Json"},
                            },
                        }
                    ],
                    "sku": {"name": "Standard"},
                    "transformation": {
                        "name": "transformationtest",
                        "properties": {"query": "Select Id, Name from inputtest", "streamingUnits": 1},
                    },
                },
                "tags": {"key1": "value1", "key3": "value3", "randomKey": "randomValue"},
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Subscription_TestQuery.json
if __name__ == "__main__":
    main()
