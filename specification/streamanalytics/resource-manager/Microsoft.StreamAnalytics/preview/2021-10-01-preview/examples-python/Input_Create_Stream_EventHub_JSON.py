from azure.identity import DefaultAzureCredential
from azure.mgmt.streamanalytics import StreamAnalyticsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-streamanalytics
# USAGE
    python input_create_stream_event_hub_json.py

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

    response = client.inputs.create_or_replace(
        resource_group_name="sjrg3139",
        job_name="sj197",
        input_name="input7425",
        input={
            "properties": {
                "datasource": {
                    "properties": {
                        "consumerGroupName": "sdkconsumergroup",
                        "eventHubName": "sdkeventhub",
                        "serviceBusNamespace": "sdktest",
                        "sharedAccessPolicyKey": "someSharedAccessPolicyKey==",
                        "sharedAccessPolicyName": "RootManageSharedAccessKey",
                    },
                    "type": "Microsoft.ServiceBus/EventHub",
                },
                "serialization": {"properties": {"encoding": "UTF8"}, "type": "Json"},
                "type": "Stream",
                "watermarkSettings": {"watermarkMode": "ReadWatermark"},
            }
        },
    )
    print(response)


# x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Input_Create_Stream_EventHub_JSON.json
if __name__ == "__main__":
    main()
