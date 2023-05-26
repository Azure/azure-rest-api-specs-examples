from azure.identity import DefaultAzureCredential
from azure.mgmt.databox import DataBoxManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-databox
# USAGE
    python jobs_patch_system_assigned_to_user_assigned.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DataBoxManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="YourSubscriptionId",
    )

    response = client.jobs.begin_update(
        resource_group_name="YourResourceGroupName",
        job_name="TestJobName1",
        job_resource_update_parameter={
            "identity": {
                "type": "SystemAssigned,UserAssigned",
                "userAssignedIdentities": {
                    "/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testIdentity": {}
                },
            },
            "properties": {
                "details": {
                    "keyEncryptionKey": {
                        "identityProperties": {
                            "type": "UserAssigned",
                            "userAssigned": {
                                "resourceId": "/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testIdentity"
                            },
                        },
                        "kekType": "CustomerManaged",
                        "kekUrl": "https://xxx.xxx.xx",
                        "kekVaultResourceID": "/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.KeyVault/vaults/YourKeyVaultName",
                    }
                }
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/JobsPatchSystemAssignedToUserAssigned.json
if __name__ == "__main__":
    main()
