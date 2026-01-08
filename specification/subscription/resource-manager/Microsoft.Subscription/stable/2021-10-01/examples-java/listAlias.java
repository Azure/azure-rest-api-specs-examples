
/**
 * Samples for Alias List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/listAlias.json
     */
    /**
     * Sample code: ListAlias.
     * 
     * @param manager Entry point to SubscriptionManager.
     */
    public static void listAlias(com.azure.resourcemanager.subscription.SubscriptionManager manager) {
        manager.alias().listWithResponse(com.azure.core.util.Context.NONE);
    }
}
