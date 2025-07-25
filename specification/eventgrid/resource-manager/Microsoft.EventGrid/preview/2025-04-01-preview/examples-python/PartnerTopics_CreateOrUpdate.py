from azure.identity import DefaultAzureCredential

from azure.mgmt.eventgrid import EventGridManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-eventgrid
# USAGE
    python partner_topics_create_or_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = EventGridManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="8f6b6269-84f2-4d09-9e31-1127efcd1e40",
    )

    response = client.partner_topics.create_or_update(
        resource_group_name="examplerg",
        partner_topic_name="examplePartnerTopicName1",
        partner_topic_info={
            "location": "westus2",
            "properties": {
                "expirationTimeIfNotActivatedUtc": "2022-03-23T23:06:13.109Z",
                "messageForActivation": "Example message for activation",
                "partnerRegistrationImmutableId": "6f541064-031d-4cc8-9ec3-a3b4fc0f7185",
                "partnerTopicFriendlyDescription": "Example description",
                "source": "ContosoCorp.Accounts.User1",
            },
        },
    )
    print(response)


# x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/PartnerTopics_CreateOrUpdate.json
if __name__ == "__main__":
    main()
