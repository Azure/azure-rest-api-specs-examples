from azure.identity import DefaultAzureCredential
from azure.mgmt.logic import LogicManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-logic
# USAGE
    python create_or_update_an_agreement.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = LogicManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="34adfa4f-cedf-4dc0-ba29-b6d1a69ab345",
    )

    response = client.integration_account_agreements.create_or_update(
        resource_group_name="testResourceGroup",
        integration_account_name="testIntegrationAccount",
        agreement_name="testAgreement",
        agreement={
            "location": "westus",
            "properties": {
                "agreementType": "AS2",
                "content": {
                    "aS2": {
                        "receiveAgreement": {
                            "protocolSettings": {
                                "acknowledgementConnectionSettings": {
                                    "ignoreCertificateNameMismatch": True,
                                    "keepHttpConnectionAlive": True,
                                    "supportHttpStatusCodeContinue": True,
                                    "unfoldHttpHeaders": True,
                                },
                                "envelopeSettings": {
                                    "autogenerateFileName": True,
                                    "fileNameTemplate": "Test",
                                    "messageContentType": "text/plain",
                                    "suspendMessageOnFileNameGenerationError": True,
                                    "transmitFileNameInMimeHeader": True,
                                },
                                "errorSettings": {"resendIfMDNNotReceived": True, "suspendDuplicateMessage": True},
                                "mdnSettings": {
                                    "dispositionNotificationTo": "http://tempuri.org",
                                    "mdnText": "Sample",
                                    "micHashingAlgorithm": "SHA1",
                                    "needMDN": True,
                                    "receiptDeliveryUrl": "http://tempuri.org",
                                    "sendInboundMDNToMessageBox": True,
                                    "sendMDNAsynchronously": True,
                                    "signMDN": True,
                                    "signOutboundMDNIfOptional": True,
                                },
                                "messageConnectionSettings": {
                                    "ignoreCertificateNameMismatch": True,
                                    "keepHttpConnectionAlive": True,
                                    "supportHttpStatusCodeContinue": True,
                                    "unfoldHttpHeaders": True,
                                },
                                "securitySettings": {
                                    "enableNRRForInboundDecodedMessages": True,
                                    "enableNRRForInboundEncodedMessages": True,
                                    "enableNRRForInboundMDN": True,
                                    "enableNRRForOutboundDecodedMessages": True,
                                    "enableNRRForOutboundEncodedMessages": True,
                                    "enableNRRForOutboundMDN": True,
                                    "overrideGroupSigningCertificate": False,
                                },
                                "validationSettings": {
                                    "checkCertificateRevocationListOnReceive": True,
                                    "checkCertificateRevocationListOnSend": True,
                                    "checkDuplicateMessage": True,
                                    "compressMessage": True,
                                    "encryptMessage": False,
                                    "encryptionAlgorithm": "AES128",
                                    "interchangeDuplicatesValidityDays": 100,
                                    "overrideMessageProperties": True,
                                    "signMessage": False,
                                },
                            },
                            "receiverBusinessIdentity": {"qualifier": "ZZ", "value": "ZZ"},
                            "senderBusinessIdentity": {"qualifier": "AA", "value": "AA"},
                        },
                        "sendAgreement": {
                            "protocolSettings": {
                                "acknowledgementConnectionSettings": {
                                    "ignoreCertificateNameMismatch": True,
                                    "keepHttpConnectionAlive": True,
                                    "supportHttpStatusCodeContinue": True,
                                    "unfoldHttpHeaders": True,
                                },
                                "envelopeSettings": {
                                    "autogenerateFileName": True,
                                    "fileNameTemplate": "Test",
                                    "messageContentType": "text/plain",
                                    "suspendMessageOnFileNameGenerationError": True,
                                    "transmitFileNameInMimeHeader": True,
                                },
                                "errorSettings": {"resendIfMDNNotReceived": True, "suspendDuplicateMessage": True},
                                "mdnSettings": {
                                    "dispositionNotificationTo": "http://tempuri.org",
                                    "mdnText": "Sample",
                                    "micHashingAlgorithm": "SHA1",
                                    "needMDN": True,
                                    "receiptDeliveryUrl": "http://tempuri.org",
                                    "sendInboundMDNToMessageBox": True,
                                    "sendMDNAsynchronously": True,
                                    "signMDN": True,
                                    "signOutboundMDNIfOptional": True,
                                },
                                "messageConnectionSettings": {
                                    "ignoreCertificateNameMismatch": True,
                                    "keepHttpConnectionAlive": True,
                                    "supportHttpStatusCodeContinue": True,
                                    "unfoldHttpHeaders": True,
                                },
                                "securitySettings": {
                                    "enableNRRForInboundDecodedMessages": True,
                                    "enableNRRForInboundEncodedMessages": True,
                                    "enableNRRForInboundMDN": True,
                                    "enableNRRForOutboundDecodedMessages": True,
                                    "enableNRRForOutboundEncodedMessages": True,
                                    "enableNRRForOutboundMDN": True,
                                    "overrideGroupSigningCertificate": False,
                                },
                                "validationSettings": {
                                    "checkCertificateRevocationListOnReceive": True,
                                    "checkCertificateRevocationListOnSend": True,
                                    "checkDuplicateMessage": True,
                                    "compressMessage": True,
                                    "encryptMessage": False,
                                    "encryptionAlgorithm": "AES128",
                                    "interchangeDuplicatesValidityDays": 100,
                                    "overrideMessageProperties": True,
                                    "signMessage": False,
                                },
                            },
                            "receiverBusinessIdentity": {"qualifier": "AA", "value": "AA"},
                            "senderBusinessIdentity": {"qualifier": "ZZ", "value": "ZZ"},
                        },
                    }
                },
                "guestIdentity": {"qualifier": "AA", "value": "AA"},
                "guestPartner": "GuestPartner",
                "hostIdentity": {"qualifier": "ZZ", "value": "ZZ"},
                "hostPartner": "HostPartner",
                "metadata": {},
            },
            "tags": {"IntegrationAccountAgreement": "<IntegrationAccountAgreementName>"},
        },
    )
    print(response)


# x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountAgreements_CreateOrUpdate.json
if __name__ == "__main__":
    main()
