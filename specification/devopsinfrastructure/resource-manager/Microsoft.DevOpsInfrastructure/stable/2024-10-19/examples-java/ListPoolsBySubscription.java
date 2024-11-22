
/**
 * Samples for Pools List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-10-19/ListPoolsBySubscription.json
     */
    /**
     * Sample code: Pools_ListBySubscription.
     * 
     * @param manager Entry point to DevOpsInfrastructureManager.
     */
    public static void
        poolsListBySubscription(com.azure.resourcemanager.devopsinfrastructure.DevOpsInfrastructureManager manager) {
        manager.pools().list(com.azure.core.util.Context.NONE);
    }
}
