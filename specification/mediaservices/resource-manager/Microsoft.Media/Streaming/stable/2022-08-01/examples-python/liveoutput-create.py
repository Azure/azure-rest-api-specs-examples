from azure.identity import DefaultAzureCredential
from azure.mgmt.media import AzureMediaServices

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-media
# USAGE
    python liveoutputcreate.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureMediaServices(
        credential=DefaultAzureCredential(),
        subscription_id="0a6ec948-5a62-437d-b9df-934dc7c1b722",
    )

    response = client.live_outputs.begin_create(
        resource_group_name="mediaresources",
        account_name="slitestmedia10",
        live_event_name="myLiveEvent1",
        live_output_name="myLiveOutput1",
        parameters={
            "properties": {
                "archiveWindowLength": "PT5M",
                "assetName": "6f3264f5-a189-48b4-a29a-a40f22575212",
                "description": "test live output 1",
                "hls": {"fragmentsPerTsSegment": 5},
                "manifestName": "testmanifest",
                "rewindWindowLength": "PT4M",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-08-01/examples/liveoutput-create.json
if __name__ == "__main__":
    main()
