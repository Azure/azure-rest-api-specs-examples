from azure.identity import DefaultAzureCredential

from azure.mgmt.network import NetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-network
# USAGE
    python azure_firewall_packet_capture.py

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

    client.azure_firewalls.begin_packet_capture(
        resource_group_name="rg1",
        azure_firewall_name="azureFirewall1",
        parameters={
            "durationInSeconds": 300,
            "fileName": "azureFirewallPacketCapture",
            "filters": [
                {"destinationPorts": ["4500"], "destinations": ["20.1.2.0"], "sources": ["20.1.1.0"]},
                {"destinationPorts": ["123", "80"], "destinations": ["10.1.2.0"], "sources": ["10.1.1.0", "10.1.1.1"]},
            ],
            "flags": [{"type": "syn"}, {"type": "fin"}],
            "numberOfPacketsToCapture": 5000,
            "protocol": "Any",
            "sasUrl": "someSASURL",
        },
    ).result()


# x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/AzureFirewallPacketCapture.json
if __name__ == "__main__":
    main()
