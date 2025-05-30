from azure.identity import DefaultAzureCredential

from azure.mgmt.eventgrid import EventGridManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-eventgrid
# USAGE
    python partner_namespaces_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = EventGridManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="5b4b650e-28b9-4790-b3ab-ddbd88d727c4",
    )

    response = client.partner_namespaces.begin_update(
        resource_group_name="examplerg",
        partner_namespace_name="examplePartnerNamespaceName1",
        partner_namespace_update_parameters={"tags": {"tag1": "value1"}},
    ).result()
    print(response)


# x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/PartnerNamespaces_Update.json
if __name__ == "__main__":
    main()
