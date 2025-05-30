from azure.identity import DefaultAzureCredential

from azure.mgmt.network import NetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-network
# USAGE
    python network_watcher_connection_monitor_create_with_arc_network.py

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

    response = client.connection_monitors.begin_create_or_update(
        resource_group_name="rg1",
        network_watcher_name="nw1",
        connection_monitor_name="cm1",
        parameters={
            "properties": {
                "endpoints": [
                    {
                        "name": "vm1",
                        "resourceId": "/subscriptions/9cece3e3-0f7d-47ca-af0e-9772773f90b7/resourceGroups/testRG/providers/Microsoft.Compute/virtualMachines/TESTVM",
                        "type": "AzureVM",
                    },
                    {"address": "bing.com", "name": "bing", "type": "ExternalAddress"},
                    {"address": "google.com", "name": "google", "type": "ExternalAddress"},
                    {
                        "locationDetails": {"region": "eastus"},
                        "name": "ArcBasedNetwork",
                        "scope": {"include": [{"address": "172.21.128.0/20"}]},
                        "subscriptionId": "9cece3e3-0f7d-47ca-af0e-9772773f90b7",
                        "type": "AzureArcNetwork",
                    },
                ],
                "outputs": [],
                "testConfigurations": [
                    {
                        "name": "testConfig1",
                        "protocol": "Tcp",
                        "tcpConfiguration": {"disableTraceRoute": False, "port": 80},
                        "testFrequencySec": 60,
                    }
                ],
                "testGroups": [
                    {
                        "destinations": ["bing", "google"],
                        "disable": False,
                        "name": "test1",
                        "sources": ["vm1", "ArcBasedNetwork"],
                        "testConfigurations": ["testConfig1"],
                    }
                ],
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NetworkWatcherConnectionMonitorCreateWithArcNetwork.json
if __name__ == "__main__":
    main()
