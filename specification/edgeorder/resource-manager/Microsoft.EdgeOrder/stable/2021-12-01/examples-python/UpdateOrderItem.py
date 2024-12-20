from azure.identity import DefaultAzureCredential

from azure.mgmt.edgeorder import EdgeOrderManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-edgeorder
# USAGE
    python update_order_item.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = EdgeOrderManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="YourSubscriptionId",
    )

    response = client.begin_update_order_item(
        order_item_name="TestOrderItemName3",
        resource_group_name="YourResourceGroupName",
        order_item_update_parameter={
            "properties": {"preferences": {"transportPreferences": {"preferredShipmentType": "CustomerManaged"}}}
        },
    ).result()
    print(response)


# x-ms-original-file: specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/UpdateOrderItem.json
if __name__ == "__main__":
    main()
