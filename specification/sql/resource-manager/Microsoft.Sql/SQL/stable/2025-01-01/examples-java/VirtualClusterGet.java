
/**
 * Samples for VirtualClusters GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/VirtualClusterGet.json
     */
    /**
     * Sample code: Gets a virtual cluster.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsAVirtualCluster(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getVirtualClusters().getByResourceGroupWithResponse("testrg",
            "vc-f769ed71-b3ad-491a-a9d5-26eeceaa6be2", com.azure.core.util.Context.NONE);
    }
}
