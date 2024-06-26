from azure.identity import DefaultAzureCredential
from azure.mgmt.labservices import ManagedLabsClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-labservices
# USAGE
    python put_lab.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ManagedLabsClient(
        credential=DefaultAzureCredential(),
        subscription_id="34adfa4f-cedf-4dc0-ba29-b6d1a69ab345",
    )

    response = client.labs.begin_create_or_update(
        resource_group_name="testrg123",
        lab_name="testlab",
        body={
            "location": "westus",
            "properties": {
                "autoShutdownProfile": {
                    "disconnectDelay": "PT5M",
                    "idleDelay": "PT5M",
                    "noConnectDelay": "PT5M",
                    "shutdownOnDisconnect": "Enabled",
                    "shutdownOnIdle": "UserAbsence",
                    "shutdownWhenNotConnected": "Enabled",
                },
                "connectionProfile": {
                    "clientRdpAccess": "Public",
                    "clientSshAccess": "Public",
                    "webRdpAccess": "None",
                    "webSshAccess": "None",
                },
                "description": "This is a test lab.",
                "labPlanId": "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.LabServices/labPlans/testlabplan",
                "networkProfile": {
                    "subnetId": "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/default"
                },
                "securityProfile": {"openAccess": "Disabled"},
                "state": "Draft",
                "title": "Test Lab",
                "virtualMachineProfile": {
                    "additionalCapabilities": {"installGpuDrivers": "Disabled"},
                    "adminUser": {"username": "test-user"},
                    "createOption": "TemplateVM",
                    "imageReference": {
                        "offer": "WindowsServer",
                        "publisher": "Microsoft",
                        "sku": "2019-Datacenter",
                        "version": "2019.0.20190410",
                    },
                    "sku": {"name": "Medium"},
                    "usageQuota": "PT10H",
                    "useSharedPassword": "Disabled",
                },
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/Labs/putLab.json
if __name__ == "__main__":
    main()
