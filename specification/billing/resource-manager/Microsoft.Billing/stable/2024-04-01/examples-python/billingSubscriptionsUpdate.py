from azure.identity import DefaultAzureCredential

from azure.mgmt.billing import BillingManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-billing
# USAGE
    python billing_subscriptions_update.py

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

    response = client.billing_subscriptions.begin_update(
        billing_account_name="00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31",
        billing_subscription_name="11111111-1111-1111-1111-111111111111",
        parameters={"properties": {"consumptionCostCenter": "ABC1234"}},
    ).result()
    print(response)


# x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingSubscriptionsUpdate.json
if __name__ == "__main__":
    main()
