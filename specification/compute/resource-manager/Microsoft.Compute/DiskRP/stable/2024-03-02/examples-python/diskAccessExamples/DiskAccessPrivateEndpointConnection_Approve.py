from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.compute import ComputeManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-compute
# USAGE
    python disk_access_private_endpoint_connection_approve.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ComputeManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="{subscription-id}",
    )

    response = client.disk_accesses.begin_update_a_private_endpoint_connection(
        resource_group_name="myResourceGroup",
        disk_access_name="myDiskAccess",
        private_endpoint_connection_name="myPrivateEndpointConnection",
        private_endpoint_connection={
            "properties": {
                "privateLinkServiceConnectionState": {
                    "description": "Approving myPrivateEndpointConnection",
                    "status": "Approved",
                }
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2024-03-02/examples/diskAccessExamples/DiskAccessPrivateEndpointConnection_Approve.json
if __name__ == "__main__":
    main()
