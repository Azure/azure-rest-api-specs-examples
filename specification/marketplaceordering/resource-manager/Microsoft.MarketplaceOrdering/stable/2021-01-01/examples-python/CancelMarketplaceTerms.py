from azure.identity import DefaultAzureCredential

from azure.mgmt.marketplaceordering import MarketplaceOrderingAgreements

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-marketplaceordering
# USAGE
    python cancel_marketplace_terms.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MarketplaceOrderingAgreements(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.marketplace_agreements.cancel(
        publisher_id="pubid",
        offer_id="offid",
        plan_id="planid",
    )
    print(response)


# x-ms-original-file: specification/marketplaceordering/resource-manager/Microsoft.MarketplaceOrdering/stable/2021-01-01/examples/CancelMarketplaceTerms.json
if __name__ == "__main__":
    main()
