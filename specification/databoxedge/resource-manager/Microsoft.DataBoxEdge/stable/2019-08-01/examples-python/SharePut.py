from azure.identity import DefaultAzureCredential

from azure.mgmt.databoxedge import DataBoxEdgeManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-databoxedge
# USAGE
    python share_put.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DataBoxEdgeManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="4385cf00-2d3a-425a-832f-f4285b1c9dce",
    )

    response = client.shares.begin_create_or_update(
        device_name="testedgedevice",
        name="smbshare",
        resource_group_name="GroupForEdgeAutomation",
        share={
            "properties": {
                "accessProtocol": "SMB",
                "azureContainerInfo": {
                    "containerName": "testContainerSMB",
                    "dataFormat": "BlockBlob",
                    "storageAccountCredentialId": "/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/storageAccountCredentials/sac1",
                },
                "dataPolicy": "Cloud",
                "description": "",
                "monitoringStatus": "Enabled",
                "shareStatus": "Online",
                "userAccessRights": [
                    {
                        "accessType": "Change",
                        "userId": "/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/users/user2",
                    }
                ],
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/SharePut.json
if __name__ == "__main__":
    main()
