
/**
 * Samples for Clusters GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Clusters_GetMinimumDefaults.json
     */
    /**
     * Sample code: Get a minimum HorizonDB cluster with materialized service defaults.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void getAMinimumHorizonDBClusterWithMaterializedServiceDefaults(
        com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.clusters().getByResourceGroupWithResponse("example-resource-group", "minimum-defaults",
            com.azure.core.util.Context.NONE);
    }
}
