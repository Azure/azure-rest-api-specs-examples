/** Samples for Alias Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/subscription/resource-manager/Microsoft.Subscription/stable/2020-09-01/examples/getAlias.json
     */
    /**
     * Sample code: GetAlias.
     *
     * @param manager Entry point to SubscriptionManager.
     */
    public static void getAlias(com.azure.resourcemanager.subscription.SubscriptionManager manager) {
        manager.alias().getWithResponse("aliasForNewSub", com.azure.core.util.Context.NONE);
    }
}
