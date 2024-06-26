from azure.identity import DefaultAzureCredential
from azure.mgmt.logic import LogicManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-logic
# USAGE
    python log_a_tracked_event.py

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

    response = client.integration_accounts.log_tracking_events(
        resource_group_name="testResourceGroup",
        integration_account_name="testIntegrationAccount",
        log_tracking_events={
            "events": [
                {
                    "error": {"code": "NotFound", "message": "Some error occurred"},
                    "eventLevel": "Informational",
                    "eventTime": "2016-08-05T01:54:49.505567Z",
                    "record": {
                        "agreementProperties": {
                            "agreementName": "testAgreement",
                            "as2From": "testas2from",
                            "as2To": "testas2to",
                            "receiverPartnerName": "testPartner2",
                            "senderPartnerName": "testPartner1",
                        },
                        "messageProperties": {
                            "IsMessageEncrypted": False,
                            "IsMessageSigned": False,
                            "correlationMessageId": "Unique message identifier",
                            "direction": "Receive",
                            "dispositionType": "received-success",
                            "fileName": "test",
                            "isMdnExpected": True,
                            "isMessageCompressed": False,
                            "isMessageFailed": False,
                            "isNrrEnabled": True,
                            "mdnType": "Async",
                            "messageId": "12345",
                        },
                    },
                    "recordType": "AS2Message",
                }
            ],
            "sourceType": "Microsoft.Logic/workflows",
        },
    )
    print(response)


# x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccounts_LogTrackingEvents.json
if __name__ == "__main__":
    main()
