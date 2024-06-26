from azure.identity import DefaultAzureCredential
from azure.mgmt.datalake.analytics import DataLakeAnalyticsAccountManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-datalake-analytics
# USAGE
    python operations_list.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DataLakeAnalyticsAccountManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.operations.list()
    print(response)


# x-ms-original-file: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/Operations_List.json
if __name__ == "__main__":
    main()
