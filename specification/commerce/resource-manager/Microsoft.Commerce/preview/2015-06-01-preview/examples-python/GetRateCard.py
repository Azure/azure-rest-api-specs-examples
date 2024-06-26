from azure.identity import DefaultAzureCredential
from azure.mgmt.commerce import UsageManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-commerce
# USAGE
    python get_rate_card.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = UsageManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="6d61cc05-8f8f-4916-b1b9-f1d9c25aae27",
    )

    response = client.rate_card.get(
        filter="OfferDurableId eq 'MS-AZR-0003P' and Currency eq 'USD' and Locale eq 'en-US' and RegionInfo eq 'US'",
    )
    print(response)


# x-ms-original-file: specification/commerce/resource-manager/Microsoft.Commerce/preview/2015-06-01-preview/examples/GetRateCard.json
if __name__ == "__main__":
    main()
