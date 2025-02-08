from azure.identity import DefaultAzureCredential

from azure.mgmt.monitor import MonitorManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-monitor
# USAGE
    python get_test_notifications_at_resource_group_level.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MonitorManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="187f412d-1758-44d9-b052-169e2564721d",
    )

    response = client.action_groups.get_test_notifications_at_resource_group_level(
        resource_group_name="Default-TestNotifications",
        notification_id="11000222191287",
    )
    print(response)


# x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-01/examples/getTestNotificationsAtResourceGroupLevel.json
if __name__ == "__main__":
    main()
