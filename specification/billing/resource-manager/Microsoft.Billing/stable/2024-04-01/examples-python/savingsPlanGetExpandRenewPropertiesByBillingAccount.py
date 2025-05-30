from azure.identity import DefaultAzureCredential

from azure.mgmt.billing import BillingManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-billing
# USAGE
    python savings_plan_get_expand_renew_properties_by_billing_account.py

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

    response = client.savings_plans.get_by_billing_account(
        billing_account_name="00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31",
        savings_plan_order_id="20000000-0000-0000-0000-000000000000",
        savings_plan_id="30000000-0000-0000-0000-000000000000",
    )
    print(response)


# x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/savingsPlanGetExpandRenewPropertiesByBillingAccount.json
if __name__ == "__main__":
    main()
