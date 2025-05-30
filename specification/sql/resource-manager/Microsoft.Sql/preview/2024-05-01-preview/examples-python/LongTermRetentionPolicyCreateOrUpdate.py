from azure.identity import DefaultAzureCredential

from azure.mgmt.sql import SqlManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-sql
# USAGE
    python long_term_retention_policy_create_or_update.py

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

    response = client.long_term_retention_policies.begin_create_or_update(
        resource_group_name="resourceGroup",
        server_name="testserver",
        database_name="testDatabase",
        policy_name="default",
        parameters={
            "properties": {
                "monthlyRetention": "P1Y",
                "weekOfYear": 5,
                "weeklyRetention": "P1M",
                "yearlyRetention": "P5Y",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2024-05-01-preview/examples/LongTermRetentionPolicyCreateOrUpdate.json
if __name__ == "__main__":
    main()
