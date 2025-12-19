
/**
 * Samples for WorkloadNetworks GetVirtualMachine.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/WorkloadNetworks_GetVirtualMachine.json
     */
    /**
     * Sample code: WorkloadNetworks_GetVirtualMachine.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksGetVirtualMachine(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().getVirtualMachineWithResponse("group1", "cloud1", "vm1",
            com.azure.core.util.Context.NONE);
    }
}
