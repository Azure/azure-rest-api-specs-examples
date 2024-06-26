from azure.identity import DefaultAzureCredential
from azure.mgmt.powerbidedicated import PowerBIDedicated

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-powerbidedicated
# USAGE
    python get_details_of_a_capacity.py

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

    response = client.capacities.get_details(
        resource_group_name="TestRG",
        dedicated_capacity_name="azsdktest",
    )
    print(response)


# x-ms-original-file: specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/getCapacity.json
if __name__ == "__main__":
    main()
