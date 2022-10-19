import com.azure.core.util.Context;

/** Samples for Clusters ListZones. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/Clusters_ListZones_Stretched.json
     */
    /**
     * Sample code: Clusters_ListZoneData_Stretched.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void clustersListZoneDataStretched(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.clusters().listZonesWithResponse("group1", "cloud1", "cluster1", Context.NONE);
    }
}
