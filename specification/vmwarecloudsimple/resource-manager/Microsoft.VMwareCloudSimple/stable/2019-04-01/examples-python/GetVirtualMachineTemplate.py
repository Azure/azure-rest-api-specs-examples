from azure.identity import DefaultAzureCredential
from azure.mgmt.vmwarecloudsimple import VMwareCloudSimple

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-vmwarecloudsimple
# USAGE
    python get_virtual_machine_template.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = VMwareCloudSimple(
        credential=DefaultAzureCredential(),
        subscription_id="{subscription-id}",
    )

    response = client.virtual_machine_templates.get(
        region_id="westus2",
        pc_name="myPrivateCloud",
        virtual_machine_template_name="vm-34",
    )
    print(response)


# x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/GetVirtualMachineTemplate.json
if __name__ == "__main__":
    main()
