from azure.identity import DefaultAzureCredential
from azure.mgmt.media import AzureMediaServices

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-media
# USAGE
    python contentkeypoliciescreateplayreadyopen.py

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
        content_key_policy_name="PolicyWithPlayReadyOptionAndOpenRestriction",
        parameters={
            "properties": {
                "description": "ArmPolicyDescription",
                "options": [
                    {
                        "configuration": {
                            "@odata.type": "#Microsoft.Media.ContentKeyPolicyPlayReadyConfiguration",
                            "licenses": [
                                {
                                    "allowTestDevices": True,
                                    "beginDate": "2017-10-16T18:22:53.46Z",
                                    "contentKeyLocation": {
                                        "@odata.type": "#Microsoft.Media.ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeader"
                                    },
                                    "contentType": "UltraVioletDownload",
                                    "licenseType": "Persistent",
                                    "playRight": {
                                        "allowPassingVideoContentToUnknownOutput": "NotAllowed",
                                        "digitalVideoOnlyContentRestriction": False,
                                        "imageConstraintForAnalogComponentVideoRestriction": True,
                                        "imageConstraintForAnalogComputerMonitorRestriction": False,
                                        "scmsRestriction": 2,
                                    },
                                    "securityLevel": "SL150",
                                }
                            ],
                        },
                        "name": "ArmPolicyOptionName",
                        "restriction": {"@odata.type": "#Microsoft.Media.ContentKeyPolicyOpenRestriction"},
                    }
                ],
            }
        },
    )
    print(response)


# x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/content-key-policies-create-playready-open.json
if __name__ == "__main__":
    main()
