from azure.identity import DefaultAzureCredential
from azure.mgmt.cosmosdbforpostgresql import CosmosdbForPostgresqlMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-cosmosdbforpostgresql
# USAGE
    python firewall_rule_delete.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = CosmosdbForPostgresqlMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="ffffffff-ffff-ffff-ffff-ffffffffffff",
    )

    client.firewall_rules.begin_delete(
        resource_group_name="TestGroup",
        cluster_name="pgtestsvc4",
        firewall_rule_name="rule1",
    ).result()


# x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/FirewallRuleDelete.json
if __name__ == "__main__":
    main()
