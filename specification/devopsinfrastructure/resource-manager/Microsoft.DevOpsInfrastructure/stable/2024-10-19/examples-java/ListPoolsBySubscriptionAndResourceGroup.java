
/**
 * Samples for Pools ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-10-19/ListPoolsBySubscriptionAndResourceGroup.json
     */
    /**
     * Sample code: Pools_ListByResourceGroup.
     * 
     * @param manager Entry point to DevOpsInfrastructureManager.
     */
    public static void
        poolsListByResourceGroup(com.azure.resourcemanager.devopsinfrastructure.DevOpsInfrastructureManager manager) {
        manager.pools().listByResourceGroup("rg", com.azure.core.util.Context.NONE);
    }
}
