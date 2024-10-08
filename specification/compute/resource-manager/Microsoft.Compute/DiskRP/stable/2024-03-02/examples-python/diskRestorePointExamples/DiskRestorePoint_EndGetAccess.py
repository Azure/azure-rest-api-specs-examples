from azure.identity import DefaultAzureCredential

from azure.mgmt.compute import ComputeManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-compute
# USAGE
    python disk_restore_point_end_get_access.py

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

    client.disk_restore_point.begin_revoke_access(
        resource_group_name="myResourceGroup",
        restore_point_collection_name="rpc",
        vm_restore_point_name="vmrp",
        disk_restore_point_name="TestDisk45ceb03433006d1baee0_b70cd924-3362-4a80-93c2-9415eaa12745",
    ).result()


# x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2024-03-02/examples/diskRestorePointExamples/DiskRestorePoint_EndGetAccess.json
if __name__ == "__main__":
    main()
