
/**
 * Samples for ResourceProvider GetOrderItemByName.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/GetOrderItemByName.json
     */
    /**
     * Sample code: GetOrderItemByName.
     * 
     * @param manager Entry point to EdgeOrderManager.
     */
    public static void getOrderItemByName(com.azure.resourcemanager.edgeorder.EdgeOrderManager manager) {
        manager.resourceProviders().getOrderItemByNameWithResponse("TestOrderItemName1", "YourResourceGroupName", null,
            com.azure.core.util.Context.NONE);
    }
}
