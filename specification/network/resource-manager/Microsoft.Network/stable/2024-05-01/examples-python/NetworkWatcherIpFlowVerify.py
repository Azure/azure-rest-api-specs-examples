from azure.identity import DefaultAzureCredential

from azure.mgmt.network import NetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-network
# USAGE
    python network_watcher_ip_flow_verify.py

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

    response = client.network_watchers.begin_verify_ip_flow(
        resource_group_name="rg1",
        network_watcher_name="nw1",
        parameters={
            "direction": "Outbound",
            "localIPAddress": "10.2.0.4",
            "localPort": "80",
            "protocol": "TCP",
            "remoteIPAddress": "121.10.1.1",
            "remotePort": "80",
            "targetResourceId": "/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1",
        },
    ).result()
    print(response)


# x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/NetworkWatcherIpFlowVerify.json
if __name__ == "__main__":
    main()
