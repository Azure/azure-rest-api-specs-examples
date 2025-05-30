from azure.identity import DefaultAzureCredential

from azure.mgmt.eventhub import EventHubManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-eventhub
# USAGE
    python eh_event_hub_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = EventHubManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="5f750a97-50d9-4e36-8081-c9ee4c0210d4",
    )

    response = client.event_hubs.create_or_update(
        resource_group_name="Default-NotificationHubs-AustraliaEast",
        namespace_name="sdk-Namespace-5357",
        event_hub_name="sdk-EventHub-6547",
        parameters={
            "properties": {
                "captureDescription": {
                    "destination": {
                        "identity": {
                            "type": "UserAssigned",
                            "userAssignedIdentity": "/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ud2",
                        },
                        "name": "EventHubArchive.AzureBlockBlob",
                        "properties": {
                            "archiveNameFormat": "{Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}",
                            "blobContainer": "container",
                            "storageAccountResourceId": "/subscriptions/e2f361f0-3b27-4503-a9cc-21cfba380093/resourceGroups/Default-Storage-SouthCentralUS/providers/Microsoft.ClassicStorage/storageAccounts/arjunteststorage",
                        },
                    },
                    "enabled": True,
                    "encoding": "Avro",
                    "intervalInSeconds": 120,
                    "sizeLimitInBytes": 10485763,
                },
                "messageRetentionInDays": 4,
                "partitionCount": 4,
                "retentionDescription": {
                    "cleanupPolicy": "Compact",
                    "retentionTimeInHours": 96,
                    "tombstoneRetentionTimeInHours": 1,
                },
                "status": "Active",
                "userMetadata": "key",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/EventHubs/EHEventHubCreate.json
if __name__ == "__main__":
    main()
