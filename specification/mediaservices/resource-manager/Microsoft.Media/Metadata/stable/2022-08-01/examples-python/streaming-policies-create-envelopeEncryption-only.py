from azure.identity import DefaultAzureCredential
from azure.mgmt.media import AzureMediaServices

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-media
# USAGE
    python streamingpoliciescreateenvelope_encryptiononly.py

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

    response = client.streaming_policies.create(
        resource_group_name="contoso",
        account_name="contosomedia",
        streaming_policy_name="UserCreatedSecureStreamingPolicyWithEnvelopeEncryptionOnly",
        parameters={
            "properties": {
                "defaultContentKeyPolicyName": "PolicyWithClearKeyOptionAndTokenRestriction",
                "envelopeEncryption": {
                    "contentKeys": {"defaultKey": {"label": "aesDefaultKey"}},
                    "customKeyAcquisitionUrlTemplate": "https://contoso.com/{AssetAlternativeId}/envelope/{ContentKeyId}",
                    "enabledProtocols": {"dash": True, "download": False, "hls": True, "smoothStreaming": True},
                },
            }
        },
    )
    print(response)


# x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-policies-create-envelopeEncryption-only.json
if __name__ == "__main__":
    main()
