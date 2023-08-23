/** Samples for WorkloadNetworks GetVMGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/WorkloadNetworks_GetVMGroups.json
     */
    /**
     * Sample code: WorkloadNetworks_GetVMGroup.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksGetVMGroup(com.azure.resourcemanager.avs.AvsManager manager) {
        manager
            .workloadNetworks()
            .getVMGroupWithResponse("group1", "cloud1", "vmGroup1", com.azure.core.util.Context.NONE);
    }
}
