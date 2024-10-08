
/**
 * Samples for ResourceProvider DeleteOrderItemByName.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/DeleteOrderItemByName.
     * json
     */
    /**
     * Sample code: DeleteOrderItemByName.
     * 
     * @param manager Entry point to EdgeOrderManager.
     */
    public static void deleteOrderItemByName(com.azure.resourcemanager.edgeorder.EdgeOrderManager manager) {
        manager.resourceProviders().deleteOrderItemByName("TestOrderItemName3", "YourResourceGroupName",
            com.azure.core.util.Context.NONE);
    }
}
