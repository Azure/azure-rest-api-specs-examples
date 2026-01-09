
/**
 * Samples for Alias Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/deleteAlias.json
     */
    /**
     * Sample code: DeleteAlias.
     * 
     * @param manager Entry point to SubscriptionManager.
     */
    public static void deleteAlias(com.azure.resourcemanager.subscription.SubscriptionManager manager) {
        manager.alias().deleteWithResponse("aliasForNewSub", com.azure.core.util.Context.NONE);
    }
}
