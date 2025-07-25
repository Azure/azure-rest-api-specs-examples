from azure.identity import DefaultAzureCredential

from azure.mgmt.eventgrid import EventGridManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-eventgrid
# USAGE
    python partner_destinations_create_or_update.py

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

    response = client.partner_destinations.begin_create_or_update(
        resource_group_name="examplerg",
        partner_destination_name="examplePartnerDestinationName1",
        partner_destination={
            "location": "westus2",
            "properties": {
                "endpointBaseUrl": "https://www.example/endpoint",
                "endpointServiceContext": "This is an example",
                "expirationTimeIfNotActivatedUtc": "2022-03-14T19:33:43.430Z",
                "messageForActivation": "Sample Activation message",
                "partnerRegistrationImmutableId": "0bd70ee2-7d95-447e-ab1f-c4f320019404",
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/PartnerDestinations_CreateOrUpdate.json
if __name__ == "__main__":
    main()
