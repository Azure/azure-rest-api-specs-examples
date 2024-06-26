from azure.identity import DefaultAzureCredential
from azure.mgmt.media import AzureMediaServices

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-media
# USAGE
    python account_filterscreate.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureMediaServices(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.account_filters.create_or_update(
        resource_group_name="contoso",
        account_name="contosomedia",
        filter_name="newAccountFilter",
        parameters={
            "properties": {
                "firstQuality": {"bitrate": 128000},
                "presentationTimeRange": {
                    "endTimestamp": 170000000,
                    "forceEndTimestamp": False,
                    "liveBackoffDuration": 0,
                    "presentationWindowDuration": 9223372036854775000,
                    "startTimestamp": 0,
                    "timescale": 10000000,
                },
                "tracks": [
                    {
                        "trackSelections": [
                            {"operation": "Equal", "property": "Type", "value": "Audio"},
                            {"operation": "NotEqual", "property": "Language", "value": "en"},
                            {"operation": "NotEqual", "property": "FourCC", "value": "EC-3"},
                        ]
                    },
                    {
                        "trackSelections": [
                            {"operation": "Equal", "property": "Type", "value": "Video"},
                            {"operation": "Equal", "property": "Bitrate", "value": "3000000-5000000"},
                        ]
                    },
                ],
            }
        },
    )
    print(response)


# x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/accountFilters-create.json
if __name__ == "__main__":
    main()
