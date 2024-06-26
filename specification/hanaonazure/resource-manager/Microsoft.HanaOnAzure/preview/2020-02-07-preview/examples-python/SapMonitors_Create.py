from azure.identity import DefaultAzureCredential
from azure.mgmt.hanaonazure import HanaManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hanaonazure
# USAGE
    python create_a_sap_monitor.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HanaManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.sap_monitors.begin_create(
        resource_group_name="myResourceGroup",
        sap_monitor_name="mySapMonitor",
        sap_monitor_parameter={
            "location": "westus",
            "properties": {
                "enableCustomerAnalytics": True,
                "logAnalyticsWorkspaceArmId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.operationalinsights/workspaces/myWorkspace",
                "logAnalyticsWorkspaceId": "00000000-0000-0000-0000-000000000000",
                "logAnalyticsWorkspaceSharedKey": "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000==",
                "monitorSubnet": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet",
            },
            "tags": {"key": "value"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/hanaonazure/resource-manager/Microsoft.HanaOnAzure/preview/2020-02-07-preview/examples/SapMonitors_Create.json
if __name__ == "__main__":
    main()
