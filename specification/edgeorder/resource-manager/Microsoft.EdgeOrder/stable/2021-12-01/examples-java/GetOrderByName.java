
/**
 * Samples for ResourceProvider GetOrderByName.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/GetOrderByName.json
     */
    /**
     * Sample code: GetOrderByName.
     * 
     * @param manager Entry point to EdgeOrderManager.
     */
    public static void getOrderByName(com.azure.resourcemanager.edgeorder.EdgeOrderManager manager) {
        manager.resourceProviders().getOrderByNameWithResponse("TestOrderName3", "YourResourceGroupName", "eastus",
            com.azure.core.util.Context.NONE);
    }
}
