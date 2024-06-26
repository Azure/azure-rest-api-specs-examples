from azure.identity import DefaultAzureCredential
from azure.mgmt.databoxedge import DataBoxEdgeManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-databoxedge
# USAGE
    python data_box_edge_device_patch.py

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

    response = client.devices.update(
        device_name="testedgedevice",
        resource_group_name="GroupForEdgeAutomation",
        parameters={
            "properties": {
                "edgeProfile": {
                    "subscription": {
                        "id": "/subscriptions/0d44739e-0563-474f-97e7-24a0cdb23b29/resourceGroups/rapvs-rg/providers/Microsoft.AzureStack/linkedSubscriptions/ca014ddc-5cf2-45f8-b390-e901e4a0ae87"
                    }
                }
            }
        },
    )
    print(response)


# x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/DataBoxEdgeDevicePatch.json
if __name__ == "__main__":
    main()
