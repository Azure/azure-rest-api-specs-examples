from azure.identity import DefaultAzureCredential
from azure.mgmt.media import AzureMediaServices

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-media
# USAGE
    python streaminglocatorscreatesecureuser_defined_content_keys.py

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

    response = client.streaming_locators.create(
        resource_group_name="contoso",
        account_name="contosomedia",
        streaming_locator_name="UserCreatedSecureStreamingLocatorWithUserDefinedContentKeys",
        parameters={
            "properties": {
                "assetName": "ClimbingMountRainier",
                "contentKeys": [
                    {
                        "id": "60000000-0000-0000-0000-000000000001",
                        "labelReferenceInStreamingPolicy": "aesDefaultKey",
                        "value": "1UqLohAfWsEGkULYxHjYZg==",
                    },
                    {
                        "id": "60000000-0000-0000-0000-000000000004",
                        "labelReferenceInStreamingPolicy": "cencDefaultKey",
                        "value": "4UqLohAfWsEGkULYxHjYZg==",
                    },
                    {
                        "id": "60000000-0000-0000-0000-000000000007",
                        "labelReferenceInStreamingPolicy": "cbcsDefaultKey",
                        "value": "7UqLohAfWsEGkULYxHjYZg==",
                    },
                ],
                "streamingLocatorId": "90000000-0000-0000-0000-00000000000A",
                "streamingPolicyName": "secureStreamingPolicy",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-locators-create-secure-userDefinedContentKeys.json
if __name__ == "__main__":
    main()
