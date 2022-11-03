from azure.identity import DefaultAzureCredential
from azure.mgmt.billing import BillingManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-billing
# USAGE
    python billing_periods_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = BillingManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.billing_periods.get(
        billing_period_name="201702-1",
    )
    print(response)


# x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/preview/2018-03-01-preview/examples/BillingPeriodsGet.json
if __name__ == "__main__":
    main()
