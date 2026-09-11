
/**
 * Samples for WorkloadNetworks DeleteVMGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/WorkloadNetworks_DeleteVMGroup.json
     */
    /**
     * Sample code: WorkloadNetworks_DeleteVMGroup.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksDeleteVMGroup(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().deleteVMGroup("group1", "vmGroup1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
