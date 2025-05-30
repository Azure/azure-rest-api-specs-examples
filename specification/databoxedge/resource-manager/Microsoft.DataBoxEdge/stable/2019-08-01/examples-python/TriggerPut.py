from azure.identity import DefaultAzureCredential

from azure.mgmt.databoxedge import DataBoxEdgeManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-databoxedge
# USAGE
    python trigger_put.py

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

    response = client.triggers.begin_create_or_update(
        device_name="testedgedevice",
        name="trigger1",
        resource_group_name="GroupForEdgeAutomation",
        trigger={
            "kind": "FileEvent",
            "properties": {
                "customContextTag": "CustomContextTags-1235346475",
                "sinkInfo": {
                    "roleId": "/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/roles/role1"
                },
                "sourceInfo": {
                    "shareId": "/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/shares/share1"
                },
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/TriggerPut.json
if __name__ == "__main__":
    main()
