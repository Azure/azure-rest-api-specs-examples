
/**
 * Samples for WorkloadNetworks ListVMGroups.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/WorkloadNetworks_ListVMGroups.json
     */
    /**
     * Sample code: WorkloadNetworks_ListVMGroups.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksListVMGroups(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().listVMGroups("group1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
