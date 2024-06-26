from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.mobilenetwork import MobileNetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-mobilenetwork
# USAGE
    python packet_core_control_plane_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MobileNetworkManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.packet_core_control_planes.begin_create_or_update(
        resource_group_name="rg1",
        packet_core_control_plane_name="TestPacketCoreCP",
        parameters={
            "location": "eastus",
            "properties": {
                "controlPlaneAccessInterface": {"name": "N2"},
                "coreNetworkTechnology": "5GC",
                "eventHub": {
                    "id": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.EventHub/namespaces/contosoNamespace/eventHubs/contosoHub",
                    "reportingInterval": 60,
                },
                "installation": {"desiredState": "Installed"},
                "localDiagnosticsAccess": {
                    "authenticationType": "AAD",
                    "httpsServerCertificate": {
                        "certificateUrl": "https://contosovault.vault.azure.net/certificates/ingress"
                    },
                },
                "platform": {
                    "azureStackEdgeDevice": {
                        "id": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/TestAzureStackEdgeDevice"
                    },
                    "connectedCluster": {
                        "id": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Kubernetes/connectedClusters/TestConnectedCluster"
                    },
                    "customLocation": {
                        "id": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ExtendedLocation/customLocations/TestCustomLocation"
                    },
                    "type": "AKS-HCI",
                },
                "signaling": {
                    "nasEncryption": ["NEA2/EEA2", "NEA1/EEA1", "NEA0/EEA0"],
                    "nasReroute": {"macroMmeGroupId": 1024},
                },
                "sites": [
                    {
                        "id": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/sites/testSite"
                    }
                ],
                "sku": "G0",
                "ueMtu": 1600,
                "version": "0.2.0",
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-04-01/examples/PacketCoreControlPlaneCreate.json
if __name__ == "__main__":
    main()
