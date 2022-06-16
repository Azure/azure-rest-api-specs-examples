import com.azure.core.util.Context;

/** Samples for ResourceProvider ListOrderItemsAtResourceGroupLevel. */
public final class Main {
    /*
     * x-ms-original-file: specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ListOrderItemsAtResourceGroupLevel.json
     */
    /**
     * Sample code: ListOrderItemsAtResourceGroupLevel.
     *
     * @param manager Entry point to EdgeOrderManager.
     */
    public static void listOrderItemsAtResourceGroupLevel(
        com.azure.resourcemanager.edgeorder.EdgeOrderManager manager) {
        manager.resourceProviders().listOrderItemsAtResourceGroupLevel("TestRG", null, null, null, Context.NONE);
    }
}
