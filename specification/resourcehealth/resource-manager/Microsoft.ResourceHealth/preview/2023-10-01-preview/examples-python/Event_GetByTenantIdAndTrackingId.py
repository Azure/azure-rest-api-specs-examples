from azure.identity import DefaultAzureCredential

from azure.mgmt.resourcehealth import ResourceHealthMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-resourcehealth
# USAGE
    python event_get_by_tenant_id_and_tracking_id.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ResourceHealthMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.event.get_by_tenant_id_and_tracking_id(
        event_tracking_id="eventTrackingId",
    )
    print(response)


# x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/Event_GetByTenantIdAndTrackingId.json
if __name__ == "__main__":
    main()
