from azure.identity import DefaultAzureCredential
from azure.mgmt.vmwarecloudsimple import VMwareCloudSimple

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-vmwarecloudsimple
# USAGE
    python list_resource_pools.py

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

    response = client.resource_pools.list(
        region_id="westus2",
        pc_name="myPrivateCloud",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/ListResourcePools.json
if __name__ == "__main__":
    main()
