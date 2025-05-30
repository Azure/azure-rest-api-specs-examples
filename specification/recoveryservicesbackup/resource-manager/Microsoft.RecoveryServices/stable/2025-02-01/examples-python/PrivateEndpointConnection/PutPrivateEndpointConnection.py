from azure.identity import DefaultAzureCredential

from azure.mgmt.recoveryservicesbackup.activestamp import RecoveryServicesBackupClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservicesbackup
# USAGE
    python put_private_endpoint_connection.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = RecoveryServicesBackupClient(
        credential=DefaultAzureCredential(),
        subscription_id="04cf684a-d41f-4550-9f70-7708a3a2283b",
    )

    response = client.private_endpoint_connection.begin_put(
        vault_name="gaallavaultbvtd2msi",
        resource_group_name="gaallaRG",
        private_endpoint_connection_name="gaallatestpe2.5704c932-249a-490b-a142-1396838cd3b",
        parameters={
            "properties": {
                "groupIds": ["AzureBackup_secondary"],
                "privateEndpoint": {
                    "id": "/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/gaallaRG/providers/Microsoft.Network/privateEndpoints/gaallatestpe3"
                },
                "privateLinkServiceConnectionState": {
                    "description": "Approved by johndoe@company.com",
                    "status": "Approved",
                },
                "provisioningState": "Succeeded",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/PrivateEndpointConnection/PutPrivateEndpointConnection.json
if __name__ == "__main__":
    main()
