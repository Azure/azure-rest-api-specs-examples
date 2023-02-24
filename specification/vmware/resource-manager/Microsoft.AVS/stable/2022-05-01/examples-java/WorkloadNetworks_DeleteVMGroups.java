/** Samples for WorkloadNetworks DeleteVMGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/WorkloadNetworks_DeleteVMGroups.json
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
