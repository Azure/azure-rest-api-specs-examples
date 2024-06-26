from azure.identity import DefaultAzureCredential
from azure.mgmt.billing import BillingManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-billing
# USAGE
    python billing_account_invoices_list_with_rebill_details.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = BillingManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.invoices.list_by_billing_account(
        billing_account_name="{billingAccountName}",
        period_start_date="2018-01-01",
        period_end_date="2018-06-30",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingAccountInvoicesListWithRebillDetails.json
if __name__ == "__main__":
    main()
