
/**
 * Samples for Pools Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Pools_Stop.json
     */
    /**
     * Sample code: Stop and deallocate a HorizonDB pool without applying pending parameter changes.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void stopAndDeallocateAHorizonDBPoolWithoutApplyingPendingParameterChanges(
        com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.pools().stop("exampleresourcegroup", "examplecluster", "example-pool-01",
            com.azure.core.util.Context.NONE);
    }
}
