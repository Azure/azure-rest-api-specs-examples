from azure.identity import DefaultAzureCredential
from azure.mgmt.storagepool import StoragePoolManagement

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-storagepool
# USAGE
    python create_or_update_i_scsi_target.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = StoragePoolManagement(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.iscsi_targets.begin_create_or_update(
        resource_group_name="myResourceGroup",
        disk_pool_name="myDiskPool",
        iscsi_target_name="myIscsiTarget",
        iscsi_target_create_payload={
            "properties": {
                "aclMode": "Dynamic",
                "luns": [
                    {
                        "managedDiskAzureResourceId": "/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vm-name_DataDisk_1",
                        "name": "lun0",
                    }
                ],
                "targetIqn": "iqn.2005-03.org.iscsi:server1",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/IscsiTargets_Put.json
if __name__ == "__main__":
    main()
