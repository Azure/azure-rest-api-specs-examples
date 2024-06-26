from azure.identity import DefaultAzureCredential
from azure.mgmt.labservices import ManagedLabsClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-labservices
# USAGE
    python put_lab_plan.py

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

    response = client.lab_plans.begin_create_or_update(
        resource_group_name="testrg123",
        lab_plan_name="testlabplan",
        body={
            "location": "westus",
            "properties": {
                "defaultAutoShutdownProfile": {
                    "disconnectDelay": "PT5M",
                    "idleDelay": "PT5M",
                    "noConnectDelay": "PT5M",
                    "shutdownOnDisconnect": "Enabled",
                    "shutdownOnIdle": "UserAbsence",
                    "shutdownWhenNotConnected": "Enabled",
                },
                "defaultConnectionProfile": {
                    "clientRdpAccess": "Public",
                    "clientSshAccess": "Public",
                    "webRdpAccess": "None",
                    "webSshAccess": "None",
                },
                "defaultNetworkProfile": {
                    "subnetId": "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/default"
                },
                "sharedGalleryId": "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Compute/galleries/testsig",
                "supportInfo": {
                    "email": "help@contoso.com",
                    "instructions": "Contact support for help.",
                    "phone": "+1-202-555-0123",
                    "url": "help.contoso.com",
                },
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/LabPlans/putLabPlan.json
if __name__ == "__main__":
    main()
