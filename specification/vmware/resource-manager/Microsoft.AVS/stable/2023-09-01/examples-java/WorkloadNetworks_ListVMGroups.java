
/**
 * Samples for WorkloadNetworks ListVMGroups.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/WorkloadNetworks_ListVMGroups.json
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
