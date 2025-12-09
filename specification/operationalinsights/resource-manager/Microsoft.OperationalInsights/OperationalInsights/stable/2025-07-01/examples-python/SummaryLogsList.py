from azure.identity import DefaultAzureCredential

from azure.mgmt.loganalytics import LogAnalyticsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-loganalytics
# USAGE
    python summary_logs_list.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = LogAnalyticsManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="11111111-1111-1111-1111-111111111111",
    )

    response = client.summary_logs.list_by_workspace(
        resource_group_name="oiautorest6685",
        workspace_name="oiautorest6685",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-07-01/examples/SummaryLogsList.json
if __name__ == "__main__":
    main()
