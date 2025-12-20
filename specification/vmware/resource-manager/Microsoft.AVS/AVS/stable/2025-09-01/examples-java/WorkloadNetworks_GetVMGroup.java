
/**
 * Samples for WorkloadNetworks GetVMGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/WorkloadNetworks_GetVMGroup.json
     */
    /**
     * Sample code: WorkloadNetworks_GetVMGroup.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksGetVMGroup(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().getVMGroupWithResponse("group1", "cloud1", "vmGroup1",
            com.azure.core.util.Context.NONE);
    }
}
