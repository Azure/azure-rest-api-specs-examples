from azure.identity import DefaultAzureCredential

from azure.mgmt.monitor import MonitorManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-monitor
# USAGE
    python get_metric_definitions_metric_class.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MonitorManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.metric_definitions.list(
        resource_uri="subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-01-01/examples/GetMetricDefinitionsMetricClass.json
if __name__ == "__main__":
    main()
