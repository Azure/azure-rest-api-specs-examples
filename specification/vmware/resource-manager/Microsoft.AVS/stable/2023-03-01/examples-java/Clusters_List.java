/** Samples for Clusters List. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/Clusters_List.json
     */
    /**
     * Sample code: Clusters_List.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void clustersList(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.clusters().list("group1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
