from azure.identity import DefaultAzureCredential
from azure.mgmt.billing import BillingManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-billing
# USAGE
    python create_billing_profile.py

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

    response = client.billing_profiles.begin_create_or_update(
        billing_account_name="{billingAccountName}",
        billing_profile_name="{billingProfileName}",
        parameters={
            "properties": {
                "billTo": {
                    "addressLine1": "Test Address 1",
                    "city": "Redmond",
                    "country": "US",
                    "firstName": "Test",
                    "lastName": "User",
                    "postalCode": "12345",
                    "region": "WA",
                },
                "displayName": "Finance",
                "enabledAzurePlans": [{"skuId": "0001"}, {"skuId": "0002"}],
                "invoiceEmailOptIn": True,
                "poNumber": "ABC12345",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/PutBillingProfile.json
if __name__ == "__main__":
    main()
