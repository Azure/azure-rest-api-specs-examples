from azure.identity import DefaultAzureCredential

from azure.mgmt.network import NetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-network
# USAGE
    python azure_firewall_put_in_hub.py

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

    response = client.azure_firewalls.begin_create_or_update(
        resource_group_name="rg1",
        azure_firewall_name="azurefirewall",
        parameters={
            "location": "West US",
            "properties": {
                "firewallPolicy": {
                    "id": "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/firewallPolicies/policy1"
                },
                "hubIPAddresses": {"publicIPs": {"addresses": [], "count": 1}},
                "sku": {"name": "AZFW_Hub", "tier": "Standard"},
                "threatIntelMode": "Alert",
                "virtualHub": {
                    "id": "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1"
                },
            },
            "tags": {"key1": "value1"},
            "zones": [],
        },
    ).result()
    print(response)


# x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/AzureFirewallPutInHub.json
if __name__ == "__main__":
    main()
