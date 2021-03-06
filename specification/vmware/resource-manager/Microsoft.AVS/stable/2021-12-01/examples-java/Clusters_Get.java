import com.azure.core.util.Context;

/** Samples for Clusters Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/Clusters_Get.json
     */
    /**
     * Sample code: Clusters_Get.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void clustersGet(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.clusters().getWithResponse("group1", "cloud1", "cluster1", Context.NONE);
    }
}
