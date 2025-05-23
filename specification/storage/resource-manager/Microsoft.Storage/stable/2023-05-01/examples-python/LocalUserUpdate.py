from azure.identity import DefaultAzureCredential

from azure.mgmt.storage import StorageManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-storage
# USAGE
    python local_user_update.py

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

    response = client.local_users.create_or_update(
        resource_group_name="res6977",
        account_name="sto2527",
        username="user1",
        properties={
            "properties": {
                "allowAclAuthorization": False,
                "extendedGroups": [1001, 1005, 2005],
                "groupId": 3000,
                "hasSharedKey": False,
                "hasSshKey": False,
                "hasSshPassword": False,
                "homeDirectory": "homedirectory2",
                "isNFSv3Enabled": True,
            }
        },
    )
    print(response)


# x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/LocalUserUpdate.json
if __name__ == "__main__":
    main()
