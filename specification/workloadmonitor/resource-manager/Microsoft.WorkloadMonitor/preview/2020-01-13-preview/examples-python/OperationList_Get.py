from azure.identity import DefaultAzureCredential
from azure.mgmt.workloadmonitor import WorkloadMonitorAPI

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-workloadmonitor
# USAGE
    python operation_list_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = WorkloadMonitorAPI(
        credential=DefaultAzureCredential(),
    )

    response = client.operations.list()
    for item in response:
        print(item)


# x-ms-original-file: specification/workloadmonitor/resource-manager/Microsoft.WorkloadMonitor/preview/2020-01-13-preview/examples/OperationList_Get.json
if __name__ == "__main__":
    main()
