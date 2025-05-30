from azure.identity import DefaultAzureCredential

from azure.mgmt.storage import StorageManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-storage
# USAGE
    python file_shares_patch_paid_bursting.py

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

    response = client.file_shares.update(
        resource_group_name="res3376",
        account_name="sto328",
        share_name="share6185",
        file_share={
            "properties": {
                "fileSharePaidBursting": {
                    "paidBurstingEnabled": True,
                    "paidBurstingMaxBandwidthMibps": 10340,
                    "paidBurstingMaxIops": 102400,
                }
            }
        },
    )
    print(response)


# x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/FileSharesPatch_PaidBursting.json
if __name__ == "__main__":
    main()
