from azure.identity import DefaultAzureCredential
from azure.mgmt.oep import OpenEnergyPlatformManagementServiceAPIs

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-oep
# USAGE
    python oep_resource_list_by_resource_group.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = OpenEnergyPlatformManagementServiceAPIs(
        credential=DefaultAzureCredential(),
        subscription_id="0000000-0000-0000-0000-000000000001",
    )

    response = client.energy_services.list_by_resource_group(
        resource_group_name="DummyResourceGroupName",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/oep/resource-manager/Microsoft.OpenEnergyPlatform/preview/2022-04-04-preview/examples/OepResource_ListByResourceGroup.json
if __name__ == "__main__":
    main()
