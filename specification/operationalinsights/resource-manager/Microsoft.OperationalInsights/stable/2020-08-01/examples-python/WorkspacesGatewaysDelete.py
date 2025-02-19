from azure.identity import DefaultAzureCredential

from azure.mgmt.loganalytics import LogAnalyticsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-loganalytics
# USAGE
    python workspaces_gateways_delete.py

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

    client.gateways.delete(
        resource_group_name="OIAutoRest5123",
        workspace_name="aztest5048",
        gateway_id="00000000-0000-0000-0000-00000000000",
    )


# x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/WorkspacesGatewaysDelete.json
if __name__ == "__main__":
    main()
