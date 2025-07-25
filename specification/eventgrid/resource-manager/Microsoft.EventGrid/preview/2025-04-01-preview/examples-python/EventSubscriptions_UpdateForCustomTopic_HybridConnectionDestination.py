from azure.identity import DefaultAzureCredential

from azure.mgmt.eventgrid import EventGridManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-eventgrid
# USAGE
    python event_subscriptions_update_for_custom_topic_hybrid_connection_destination.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = EventGridManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.event_subscriptions.begin_update(
        scope="subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic2",
        event_subscription_name="examplesubscription1",
        event_subscription_update_parameters={
            "destination": {
                "endpointType": "HybridConnection",
                "properties": {
                    "resourceId": "/subscriptions/d33c5f7a-02ea-40f4-bf52-07f17e84d6a8/resourceGroups/TestRG/providers/Microsoft.Relay/namespaces/ContosoNamespace/hybridConnections/HC1"
                },
            },
            "filter": {
                "isSubjectCaseSensitive": True,
                "subjectBeginsWith": "existingPrefix",
                "subjectEndsWith": "newSuffix",
            },
            "labels": ["label1", "label2"],
        },
    ).result()
    print(response)


# x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/EventSubscriptions_UpdateForCustomTopic_HybridConnectionDestination.json
if __name__ == "__main__":
    main()
