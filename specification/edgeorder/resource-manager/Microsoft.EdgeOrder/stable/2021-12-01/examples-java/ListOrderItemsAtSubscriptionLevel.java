import com.azure.core.util.Context;

/** Samples for ResourceProvider ListOrderItemsAtSubscriptionLevel. */
public final class Main {
    /*
     * x-ms-original-file: specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ListOrderItemsAtSubscriptionLevel.json
     */
    /**
     * Sample code: ListOrderItemsAtSubscriptionLevel.
     *
     * @param manager Entry point to EdgeOrderManager.
     */
    public static void listOrderItemsAtSubscriptionLevel(com.azure.resourcemanager.edgeorder.EdgeOrderManager manager) {
        manager.resourceProviders().listOrderItemsAtSubscriptionLevel(null, null, null, Context.NONE);
    }
}
