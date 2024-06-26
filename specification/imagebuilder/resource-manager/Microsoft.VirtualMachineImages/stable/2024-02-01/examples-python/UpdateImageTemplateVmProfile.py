from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.imagebuilder import ImageBuilderClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-imagebuilder
# USAGE
    python update_image_template_vm_profile.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ImageBuilderClient(
        credential=DefaultAzureCredential(),
        subscription_id="{subscription-id}",
    )

    response = client.virtual_machine_image_templates.begin_update(
        resource_group_name="myResourceGroup",
        image_template_name="myImageTemplate",
        parameters={
            "properties": {
                "vmProfile": {
                    "osDiskSizeGB": 127,
                    "vmSize": "{updated_vmsize}",
                    "vnetConfig": {
                        "containerInstanceSubnetId": "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/subnetname",
                        "subnetId": "{updated_aci_subnet}",
                    },
                }
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2024-02-01/examples/UpdateImageTemplateVmProfile.json
if __name__ == "__main__":
    main()
