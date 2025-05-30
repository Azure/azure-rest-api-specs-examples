from azure.identity import DefaultAzureCredential

from azure.mgmt.sql import SqlManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-sql
# USAGE
    python virtual_cluster_list_by_resource_group.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SqlManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="20d7082a-0fc7-4468-82bd-542694d5042b",
    )

    response = client.virtual_clusters.list_by_resource_group(
        resource_group_name="testrg",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2024-05-01-preview/examples/VirtualClusterListByResourceGroup.json
if __name__ == "__main__":
    main()
