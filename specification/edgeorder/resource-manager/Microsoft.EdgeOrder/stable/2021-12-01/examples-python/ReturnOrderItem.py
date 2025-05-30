from azure.identity import DefaultAzureCredential

from azure.mgmt.edgeorder import EdgeOrderManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-edgeorder
# USAGE
    python return_order_item.py

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

    client.begin_return_order_item(
        order_item_name="TestOrderName4",
        resource_group_name="YourResourceGroupName",
        return_order_item_details={"returnReason": "Order returned"},
    ).result()


# x-ms-original-file: specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ReturnOrderItem.json
if __name__ == "__main__":
    main()
