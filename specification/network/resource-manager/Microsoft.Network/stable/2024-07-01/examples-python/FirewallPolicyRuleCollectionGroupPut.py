from azure.identity import DefaultAzureCredential

from azure.mgmt.network import NetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-network
# USAGE
    python firewall_policy_rule_collection_group_put.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = NetworkManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.firewall_policy_rule_collection_groups.begin_create_or_update(
        resource_group_name="rg1",
        firewall_policy_name="firewallPolicy",
        rule_collection_group_name="ruleCollectionGroup1",
        parameters={
            "properties": {
                "priority": 100,
                "ruleCollections": [
                    {
                        "action": {"type": "Deny"},
                        "name": "Example-Filter-Rule-Collection",
                        "priority": 100,
                        "ruleCollectionType": "FirewallPolicyFilterRuleCollection",
                        "rules": [
                            {
                                "destinationAddresses": ["*"],
                                "destinationPorts": ["*"],
                                "ipProtocols": ["TCP"],
                                "name": "network-rule1",
                                "ruleType": "NetworkRule",
                                "sourceAddresses": ["10.1.25.0/24"],
                            }
                        ],
                    }
                ],
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/FirewallPolicyRuleCollectionGroupPut.json
if __name__ == "__main__":
    main()
