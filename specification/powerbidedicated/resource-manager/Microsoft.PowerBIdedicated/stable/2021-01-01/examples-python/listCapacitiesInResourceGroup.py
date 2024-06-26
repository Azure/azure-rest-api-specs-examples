from azure.identity import DefaultAzureCredential
from azure.mgmt.powerbidedicated import PowerBIDedicated

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-powerbidedicated
# USAGE
    python list_capacities_in_resource_group.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = PowerBIDedicated(
        credential=DefaultAzureCredential(),
        subscription_id="613192d7-503f-477a-9cfe-4efc3ee2bd60",
    )

    response = client.capacities.list_by_resource_group(
        resource_group_name="TestRG",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/listCapacitiesInResourceGroup.json
if __name__ == "__main__":
    main()
