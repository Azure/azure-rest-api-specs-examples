from azure.identity import DefaultAzureCredential

from azure.mgmt.monitor import MonitorManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-monitor
# USAGE
    python enable_receiver.py

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

    client.action_groups.enable_receiver(
        resource_group_name="Default-NotificationRules",
        action_group_name="SampleActionGroup",
        enable_request={"receiverName": "John Doe's mobile"},
    )


# x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-01/examples/enableReceiver.json
if __name__ == "__main__":
    main()
