from azure.identity import DefaultAzureCredential

from azure.mgmt.applicationinsights import ApplicationInsightsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-applicationinsights
# USAGE
    python deleted_workbooks_list.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ApplicationInsightsManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-1111-2222-3333-444444444444",
    )

    response = client.deleted_workbooks.list_by_subscription()
    for item in response:
        print(item)


# x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/preview/2024-02-01-preview/examples/DeletedWorkbooksList.json
if __name__ == "__main__":
    main()
