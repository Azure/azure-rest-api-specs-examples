from azure.identity import DefaultAzureCredential
from azure.mgmt.workloadmonitor import WorkloadMonitorAPI

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-workloadmonitor
# USAGE
    python get_monitor_state_change_with_expanded_values.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = WorkloadMonitorAPI(
        credential=DefaultAzureCredential(),
    )

    response = client.health_monitors.get_state_change(
        subscription_id="bc27da3b-3ba2-4e00-a6ec-1fde64aa1e21",
        resource_group_name="tugamidiAlerts",
        provider_name="Microsoft.Compute",
        resource_collection_name="virtualMachines",
        resource_name="linuxEUS",
        monitor_id="logical-disks|C@3A",
        timestamp_unix="1584316800",
    )
    print(response)


# x-ms-original-file: specification/workloadmonitor/resource-manager/Microsoft.WorkloadMonitor/preview/2020-01-13-preview/examples/MonitorStateChange_GetExpanded.json
if __name__ == "__main__":
    main()
