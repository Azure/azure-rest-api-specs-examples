from azure.identity import DefaultAzureCredential

from azure.mgmt.eventhub import EventHubManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-eventhub
# USAGE
    python eh_name_space_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = EventHubManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="SampleSubscription",
    )

    response = client.namespaces.begin_create_or_update(
        resource_group_name="ResurceGroupSample",
        namespace_name="NamespaceSample",
        parameters={
            "identity": {
                "type": "SystemAssigned, UserAssigned",
                "userAssignedIdentities": {
                    "/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ud1": {},
                    "/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ud2": {},
                },
            },
            "location": "East US",
            "properties": {
                "clusterArmId": "/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.EventHub/clusters/enc-test",
                "encryption": {
                    "keySource": "Microsoft.KeyVault",
                    "keyVaultProperties": [
                        {
                            "identity": {
                                "userAssignedIdentity": "/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ud1"
                            },
                            "keyName": "Samplekey",
                            "keyVaultUri": "https://aprao-keyvault-user.vault-int.azure-int.net/",
                        }
                    ],
                },
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-01-01-preview/examples/NameSpaces/EHNameSpaceCreate.json
if __name__ == "__main__":
    main()
