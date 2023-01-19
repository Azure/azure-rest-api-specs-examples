/** Samples for Tenants List. */
public final class Main {
    /*
     * x-ms-original-file: specification/subscription/resource-manager/Microsoft.Subscription/stable/2016-06-01/examples/listTenants.json
     */
    /**
     * Sample code: listTenants.
     *
     * @param manager Entry point to SubscriptionManager.
     */
    public static void listTenants(com.azure.resourcemanager.subscription.SubscriptionManager manager) {
        manager.tenants().list(com.azure.core.util.Context.NONE);
    }
}
