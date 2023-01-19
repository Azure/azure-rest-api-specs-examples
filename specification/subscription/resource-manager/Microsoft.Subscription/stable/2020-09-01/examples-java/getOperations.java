/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/subscription/resource-manager/Microsoft.Subscription/stable/2020-09-01/examples/getOperations.json
     */
    /**
     * Sample code: getOperations.
     *
     * @param manager Entry point to SubscriptionManager.
     */
    public static void getOperations(com.azure.resourcemanager.subscription.SubscriptionManager manager) {
        manager.operations().listWithResponse(com.azure.core.util.Context.NONE);
    }
}
