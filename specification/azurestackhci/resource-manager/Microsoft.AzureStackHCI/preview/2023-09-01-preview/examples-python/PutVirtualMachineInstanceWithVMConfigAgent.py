from azure.identity import DefaultAzureCredential
from azure.mgmt.azurestackhci import AzureStackHCIClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-azurestackhci
# USAGE
    python put_virtual_machine_instance_with_vm_config_agent.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureStackHCIClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.virtual_machine_instances.begin_create_or_update(
        resource_uri="subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/Microsoft.HybridCompute/machines/DemoVM",
        virtual_machine_instance={
            "extendedLocation": {
                "name": "/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location",
                "type": "CustomLocation",
            },
            "properties": {
                "hardwareProfile": {"vmSize": "Default"},
                "networkProfile": {"networkInterfaces": [{"id": "test-nic"}]},
                "osProfile": {
                    "adminPassword": "password",
                    "adminUsername": "localadmin",
                    "computerName": "luamaster",
                    "windowsConfiguration": {"provisionVMConfigAgent": True},
                },
                "securityProfile": {"enableTPM": True, "uefiSettings": {"secureBootEnabled": True}},
                "storageProfile": {
                    "imageReference": {
                        "id": "/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/galleryImages/test-gallery-image"
                    },
                    "vmConfigStoragePathId": "/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/storageContainers/test-container",
                },
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/PutVirtualMachineInstanceWithVMConfigAgent.json
if __name__ == "__main__":
    main()
