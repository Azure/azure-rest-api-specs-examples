from azure.identity import DefaultAzureCredential

from azure.mgmt.monitor import MonitorManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-monitor
# USAGE
    python operations_get.py

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

    response = client.monitor_operations.list()
    for item in response:
        print(item)


# x-ms-original-file: specification/monitor/resource-manager/Microsoft.Monitor/stable/2023-04-03/examples/OperationsGet.json
if __name__ == "__main__":
    main()
