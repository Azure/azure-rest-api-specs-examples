from azure.identity import DefaultAzureCredential

from azure.mgmt.billing import BillingManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-billing
# USAGE
    python check_access_by_billing_profile.py

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

    response = client.billing_permissions.check_access_by_billing_profile(
        billing_account_name="10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31",
        billing_profile_name="xxxx-xxxx-xxx-xxx",
        parameters={
            "actions": ["Microsoft.Billing/billingAccounts/read", "Microsoft.Subscription/subscriptions/write"]
        },
    )
    print(response)


# x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/checkAccessByBillingProfile.json
if __name__ == "__main__":
    main()
