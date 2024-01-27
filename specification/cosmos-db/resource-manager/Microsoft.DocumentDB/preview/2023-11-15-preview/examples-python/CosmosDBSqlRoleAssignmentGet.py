from azure.identity import DefaultAzureCredential
from azure.mgmt.cosmosdb import CosmosDBManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-cosmosdb
# USAGE
    python cosmos_db_sql_role_assignment_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = CosmosDBManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="mySubscriptionId",
    )

    response = client.sql_resources.get_sql_role_assignment(
        role_assignment_id="myRoleAssignmentId",
        resource_group_name="myResourceGroupName",
        account_name="myAccountName",
    )
    print(response)


# x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-11-15-preview/examples/CosmosDBSqlRoleAssignmentGet.json
if __name__ == "__main__":
    main()
