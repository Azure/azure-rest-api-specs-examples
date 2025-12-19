
/**
 * Samples for WorkloadNetworks List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/WorkloadNetworks_List.json
     */
    /**
     * Sample code: WorkloadNetworks_List.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksList(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().list("group1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
