from azure.identity import DefaultAzureCredential
from azure.mgmt.securityinsight import SecurityInsights

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-securityinsight
# USAGE
    python create_watchlist_item.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SecurityInsights(
        credential=DefaultAzureCredential(),
        subscription_id="d0cfe6b2-9ac0-4464-9919-dccaee2e48c0",
    )

    response = client.watchlist_items.create_or_update(
        resource_group_name="myRg",
        workspace_name="myWorkspace",
        watchlist_alias="highValueAsset",
        watchlist_item_id="82ba292c-dc97-4dfc-969d-d4dd9e666842",
        watchlist_item={
            "etag": "0300bf09-0000-0000-0000-5c37296e0000",
            "properties": {
                "itemsKeyValue": {
                    "Business tier": "10.0.2.0/24",
                    "Data tier": "10.0.2.0/24",
                    "Gateway subnet": "10.0.255.224/27",
                    "Private DMZ in": "10.0.0.0/27",
                    "Public DMZ out": "10.0.0.96/27",
                    "Web Tier": "10.0.1.0/24",
                }
            },
        },
    )
    print(response)


# x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-12-01-preview/examples/watchlists/CreateWatchlistItem.json
if __name__ == "__main__":
    main()
