from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.rdbms.postgresql import PostgreSQLManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-rdbms
# USAGE
    python virtual_network_rules_create_or_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = PostgreSQLManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="ffffffff-ffff-ffff-ffff-ffffffffffff",
    )

    response = client.virtual_network_rules.begin_create_or_update(
        resource_group_name="TestGroup",
        server_name="vnet-test-svr",
        virtual_network_rule_name="vnet-firewall-rule",
        parameters={
            "properties": {
                "ignoreMissingVnetServiceEndpoint": False,
                "virtualNetworkSubnetId": "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.Network/virtualNetworks/testvnet/subnets/testsubnet",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/VirtualNetworkRulesCreateOrUpdate.json
if __name__ == "__main__":
    main()
