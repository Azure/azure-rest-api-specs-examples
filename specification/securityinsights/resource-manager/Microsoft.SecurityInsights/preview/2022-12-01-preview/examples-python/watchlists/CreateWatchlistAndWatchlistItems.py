from azure.identity import DefaultAzureCredential
from azure.mgmt.securityinsight import SecurityInsights

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-securityinsight
# USAGE
    python create_watchlist_and_watchlist_items.py

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

    response = client.watchlists.create_or_update(
        resource_group_name="myRg",
        workspace_name="myWorkspace",
        watchlist_alias="highValueAsset",
        watchlist={
            "etag": '"0300bf09-0000-0000-0000-5c37296e0000"',
            "properties": {
                "contentType": "text/csv",
                "description": "Watchlist from CSV content",
                "displayName": "High Value Assets Watchlist",
                "itemsSearchKey": "header1",
                "numberOfLinesToSkip": 1,
                "provider": "Microsoft",
                "rawContent": "This line will be skipped\nheader1,header2\nvalue1,value2",
                "source": "watchlist.csv",
                "sourceType": "Local file",
            },
        },
    )
    print(response)


# x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-12-01-preview/examples/watchlists/CreateWatchlistAndWatchlistItems.json
if __name__ == "__main__":
    main()
