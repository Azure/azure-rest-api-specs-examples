from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.mobilenetwork import MobileNetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-mobilenetwork
# USAGE
    python packet_capture_create.py

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

    response = client.packet_captures.begin_create_or_update(
        resource_group_name="rg1",
        packet_core_control_plane_name="TestPacketCoreCP",
        packet_capture_name="pc1",
        parameters={
            "properties": {
                "bytesToCapturePerPacket": 10000,
                "networkInterfaces": [
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP",
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP",
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP/attachedDataNetworks/TestADN",
                ],
                "timeLimitInSeconds": 100,
                "totalBytesPerSession": 100000,
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-04-01/examples/PacketCaptureCreate.json
if __name__ == "__main__":
    main()
