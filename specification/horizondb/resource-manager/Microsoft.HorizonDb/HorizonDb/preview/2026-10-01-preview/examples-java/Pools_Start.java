
/**
 * Samples for Pools Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Pools_Start.json
     */
    /**
     * Sample code: Start a HorizonDB pool without applying pending parameter changes.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void startAHorizonDBPoolWithoutApplyingPendingParameterChanges(
        com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.pools().start("exampleresourcegroup", "examplecluster", "example-pool-01",
            com.azure.core.util.Context.NONE);
    }
}
