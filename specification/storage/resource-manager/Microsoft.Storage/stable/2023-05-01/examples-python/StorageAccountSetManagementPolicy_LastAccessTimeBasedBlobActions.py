from azure.identity import DefaultAzureCredential

from azure.mgmt.storage import StorageManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-storage
# USAGE
    python storage_account_set_management_policy_last_access_time_based_blob_actions.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = StorageManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="{subscription-id}",
    )

    response = client.management_policies.create_or_update(
        resource_group_name="res7687",
        account_name="sto9699",
        management_policy_name="default",
        properties={
            "properties": {
                "policy": {
                    "rules": [
                        {
                            "definition": {
                                "actions": {
                                    "baseBlob": {
                                        "delete": {"daysAfterLastAccessTimeGreaterThan": 1000},
                                        "enableAutoTierToHotFromCool": True,
                                        "tierToArchive": {"daysAfterLastAccessTimeGreaterThan": 90},
                                        "tierToCool": {"daysAfterLastAccessTimeGreaterThan": 30},
                                    },
                                    "snapshot": {"delete": {"daysAfterCreationGreaterThan": 30}},
                                },
                                "filters": {"blobTypes": ["blockBlob"], "prefixMatch": ["olcmtestcontainer"]},
                            },
                            "enabled": True,
                            "name": "olcmtest",
                            "type": "Lifecycle",
                        }
                    ]
                }
            }
        },
    )
    print(response)


# x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/StorageAccountSetManagementPolicy_LastAccessTimeBasedBlobActions.json
if __name__ == "__main__":
    main()
