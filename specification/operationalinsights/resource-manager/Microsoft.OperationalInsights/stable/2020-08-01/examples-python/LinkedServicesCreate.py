from azure.identity import DefaultAzureCredential

from azure.mgmt.loganalytics import LogAnalyticsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-loganalytics
# USAGE
    python linked_services_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = LogAnalyticsManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-00000000000",
    )

    response = client.linked_services.begin_create_or_update(
        resource_group_name="mms-eus",
        workspace_name="TestLinkWS",
        linked_service_name="Cluster",
        parameters={
            "properties": {
                "writeAccessResourceId": "/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/mms-eus/providers/Microsoft.OperationalInsights/clusters/testcluster"
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/LinkedServicesCreate.json
if __name__ == "__main__":
    main()
