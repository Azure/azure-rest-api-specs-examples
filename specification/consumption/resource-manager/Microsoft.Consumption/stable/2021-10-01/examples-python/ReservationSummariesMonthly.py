from azure.identity import DefaultAzureCredential
from azure.mgmt.consumption import ConsumptionManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-consumption
# USAGE
    python reservation_summaries_monthly.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ConsumptionManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.reservations_summaries.list_by_reservation_order(
        reservation_order_id="00000000-0000-0000-0000-000000000000",
        grain="monthly",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationSummariesMonthly.json
if __name__ == "__main__":
    main()
