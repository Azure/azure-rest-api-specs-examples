from azure.identity import DefaultAzureCredential
from azure.mgmt.reservations import AzureReservationAPI

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-reservations
# USAGE
    python get_compute_one_sku_usages.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureReservationAPI(
        credential=DefaultAzureCredential(),
    )

    response = client.quota.get(
        subscription_id="00000000-0000-0000-0000-000000000000",
        provider_id="Microsoft.Compute",
        location="eastus",
        resource_name="standardNDSFamily",
    )
    print(response)


# x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2020-10-25/examples/getComputeOneSkuUsages.json
if __name__ == "__main__":
    main()
