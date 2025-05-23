from azure.identity import DefaultAzureCredential

from azure.mgmt.computefleet import ComputeFleetMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-computefleet
# USAGE
    python fleets_list_virtual_machine_scale_sets.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ComputeFleetMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.fleets.list_virtual_machine_scale_sets(
        resource_group_name="rgazurefleet",
        name="myFleet",
    )
    for item in response:
        print(item)


# x-ms-original-file: 2024-11-01/Fleets_ListVirtualMachineScaleSets.json
if __name__ == "__main__":
    main()
