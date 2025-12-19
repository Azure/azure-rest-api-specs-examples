
/**
 * Samples for WorkloadNetworks Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/WorkloadNetworks_Get.json
     */
    /**
     * Sample code: WorkloadNetworks_Get.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksGet(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().getWithResponse("group1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
