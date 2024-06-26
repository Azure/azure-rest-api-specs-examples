from azure.identity import DefaultAzureCredential
from azure.mgmt.oep import OpenEnergyPlatformManagementServiceAPIs

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-oep
# USAGE
    python energy_services_remove_partition_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = OpenEnergyPlatformManagementServiceAPIs(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.energy_services.begin_remove_partition(
        resource_group_name="rgoep",
        resource_name="aaaaaaa",
    ).result()
    print(response)


# x-ms-original-file: specification/oep/resource-manager/Microsoft.OpenEnergyPlatform/preview/2022-04-04-preview/examples/EnergyServices_RemovePartition_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
