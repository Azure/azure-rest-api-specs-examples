from azure.identity import DefaultAzureCredential
from azure.mgmt.azurestackhci import AzureStackHCIClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-azurestackhci
# USAGE
    python put_virtual_hard_disk.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureStackHCIClient(
        credential=DefaultAzureCredential(),
        subscription_id="fd3c3665-1729-4b7b-9a38-238e83b0f98b",
    )

    response = client.virtual_hard_disks.begin_create_or_update(
        resource_group_name="test-rg",
        virtual_hard_disk_name="test-vhd",
        virtual_hard_disks={
            "extendedLocation": {
                "name": "/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location",
                "type": "CustomLocation",
            },
            "location": "West US2",
            "properties": {"diskSizeGB": 32},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/PutVirtualHardDisk.json
if __name__ == "__main__":
    main()
