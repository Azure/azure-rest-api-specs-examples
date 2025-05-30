from azure.identity import DefaultAzureCredential

from azure.mgmt.sql import SqlManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-sql
# USAGE
    python sync_member_delete.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SqlManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-1111-2222-3333-444444444444",
    )

    client.sync_members.begin_delete(
        resource_group_name="syncgroupcrud-65440",
        server_name="syncgroupcrud-8475",
        database_name="syncgroupcrud-4328",
        sync_group_name="syncgroupcrud-3187",
        sync_member_name="syncgroupcrud-4879",
    ).result()


# x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2024-05-01-preview/examples/SyncMemberDelete.json
if __name__ == "__main__":
    main()
