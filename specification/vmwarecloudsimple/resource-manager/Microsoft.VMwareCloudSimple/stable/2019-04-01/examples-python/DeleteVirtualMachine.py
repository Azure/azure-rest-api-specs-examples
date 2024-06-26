from azure.identity import DefaultAzureCredential
from azure.mgmt.vmwarecloudsimple import VMwareCloudSimple

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-vmwarecloudsimple
# USAGE
    python delete_virtual_machine.py

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

    response = client.virtual_machines.begin_delete(
        resource_group_name="myResourceGroup",
        referer="https://management.azure.com/",
        virtual_machine_name="myVirtualMachine",
    ).result()
    print(response)


# x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/DeleteVirtualMachine.json
if __name__ == "__main__":
    main()
