
/**
 * Samples for Clusters ListZones.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Clusters_ListZones.json
     */
    /**
     * Sample code: Clusters_ListZones.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void clustersListZones(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.clusters().listZonesWithResponse("group1", "cloud1", "cluster1", com.azure.core.util.Context.NONE);
    }
}
