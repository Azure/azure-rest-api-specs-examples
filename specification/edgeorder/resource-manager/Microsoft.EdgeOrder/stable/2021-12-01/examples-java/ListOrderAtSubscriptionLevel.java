import com.azure.core.util.Context;

/** Samples for ResourceProvider ListOrderAtSubscriptionLevel. */
public final class Main {
    /*
     * x-ms-original-file: specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ListOrderAtSubscriptionLevel.json
     */
    /**
     * Sample code: ListOrderAtSubscriptionLevel.
     *
     * @param manager Entry point to EdgeOrderManager.
     */
    public static void listOrderAtSubscriptionLevel(com.azure.resourcemanager.edgeorder.EdgeOrderManager manager) {
        manager.resourceProviders().listOrderAtSubscriptionLevel(null, Context.NONE);
    }
}
