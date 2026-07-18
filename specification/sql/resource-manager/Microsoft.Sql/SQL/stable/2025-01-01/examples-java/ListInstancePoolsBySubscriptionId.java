
/**
 * Samples for InstancePools List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ListInstancePoolsBySubscriptionId.json
     */
    /**
     * Sample code: List instance pools in the subscription.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listInstancePoolsInTheSubscription(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getInstancePools().list(com.azure.core.util.Context.NONE);
    }
}
