import com.azure.core.util.Context;

/** Samples for ResourceProvider List. */
public final class Main {
    /*
     * x-ms-original-file: specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ListAddressesAtSubscriptionLevel.json
     */
    /**
     * Sample code: ListAddressesAtSubscriptionLevel.
     *
     * @param manager Entry point to EdgeOrderManager.
     */
    public static void listAddressesAtSubscriptionLevel(com.azure.resourcemanager.edgeorder.EdgeOrderManager manager) {
        manager.resourceProviders().list(null, null, Context.NONE);
    }
}
