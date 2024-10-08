from azure.identity import DefaultAzureCredential

from azure.mgmt.billing import BillingManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-billing
# USAGE
    python billing_profiles_create_or_update.py

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
        billing_account_name="00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31",
        billing_profile_name="xxxx-xxxx-xxx-xxx",
        parameters={
            "properties": {
                "billTo": {
                    "addressLine1": "Test Address1",
                    "addressLine2": "Test Address2",
                    "addressLine3": "Test Address3",
                    "city": "City",
                    "companyName": "Contoso",
                    "country": "US",
                    "email": "abc@contoso.com",
                    "firstName": "Test",
                    "isValidAddress": True,
                    "lastName": "User",
                    "phoneNumber": "000-000-0000",
                    "postalCode": "00000",
                    "region": "WA",
                },
                "displayName": "Billing Profile 1",
                "enabledAzurePlans": [{"skuId": "0001"}, {"skuId": "0002"}],
                "invoiceEmailOptIn": True,
                "poNumber": "ABC12345",
                "shipTo": {
                    "addressLine1": "Test Address1",
                    "addressLine2": "Test Address2",
                    "addressLine3": "Test Address3",
                    "city": "City",
                    "companyName": "Contoso",
                    "country": "US",
                    "email": "abc@contoso.com",
                    "firstName": "Test",
                    "isValidAddress": True,
                    "lastName": "User",
                    "phoneNumber": "000-000-0000",
                    "postalCode": "00000",
                    "region": "WA",
                },
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingProfilesCreateOrUpdate.json
if __name__ == "__main__":
    main()
