from azure.identity import DefaultAzureCredential
from azure.mgmt.mobilenetwork import MobileNetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-mobilenetwork
# USAGE
    python packet_core_control_plane_collect_diagnostics_package.py

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

    response = client.packet_core_control_planes.begin_collect_diagnostics_package(
        resource_group_name="rg1",
        packet_core_control_plane_name="TestPacketCoreCP",
        parameters={
            "storageAccountBlobUrl": "https://contosoaccount.blob.core.windows.net/container/diagnosticsPackage.zip"
        },
    ).result()
    print(response)


# x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/PacketCoreControlPlaneCollectDiagnosticsPackage.json
if __name__ == "__main__":
    main()
