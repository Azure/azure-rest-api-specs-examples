from azure.identity import DefaultAzureCredential
from azure.mgmt.synapse import SynapseManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-synapse
# USAGE
    python create_ip_firewall_rule.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SynapseManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="01234567-89ab-4def-0123-456789abcdef",
    )

    response = client.ip_firewall_rules.begin_create_or_update(
        resource_group_name="ExampleResourceGroup",
        workspace_name="ExampleWorkspace",
        rule_name="ExampleIpFirewallRule",
        ip_firewall_rule_info={"properties": {"endIpAddress": "10.0.0.254", "startIpAddress": "10.0.0.0"}},
    ).result()
    print(response)


# x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateIpFirewallRule.json
if __name__ == "__main__":
    main()
