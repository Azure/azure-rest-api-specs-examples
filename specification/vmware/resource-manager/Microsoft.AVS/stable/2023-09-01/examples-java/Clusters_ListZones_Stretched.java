
/**
 * Samples for Clusters ListZones.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/Clusters_ListZones_Stretched.json
     */
    /**
     * Sample code: Clusters_ListZones_Stretched.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void clustersListZonesStretched(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.clusters().listZonesWithResponse("group1", "cloud1", "cluster1", com.azure.core.util.Context.NONE);
    }
}
