from azure.identity import DefaultAzureCredential
from azure.mgmt.media import AzureMediaServices

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-media
# USAGE
    python contentkeypoliciescreatenodrmtoken.py

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

    response = client.content_key_policies.create_or_update(
        resource_group_name="contoso",
        account_name="contosomedia",
        content_key_policy_name="PolicyWithClearKeyOptionAndSwtTokenRestriction",
        parameters={
            "properties": {
                "description": "ArmPolicyDescription",
                "options": [
                    {
                        "configuration": {"@odata.type": "#Microsoft.Media.ContentKeyPolicyClearKeyConfiguration"},
                        "name": "ClearKeyOption",
                        "restriction": {
                            "@odata.type": "#Microsoft.Media.ContentKeyPolicyTokenRestriction",
                            "audience": "urn:audience",
                            "issuer": "urn:issuer",
                            "primaryVerificationKey": {
                                "@odata.type": "#Microsoft.Media.ContentKeyPolicySymmetricTokenKey",
                                "keyValue": "AAAAAAAAAAAAAAAAAAAAAA==",
                            },
                            "restrictionTokenType": "Swt",
                        },
                    }
                ],
            }
        },
    )
    print(response)


# x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/content-key-policies-create-nodrm-token.json
if __name__ == "__main__":
    main()
