import com.azure.core.util.Context;

/** Samples for WorkloadNetworks GetVirtualMachine. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/WorkloadNetworks_GetVirtualMachine.json
     */
    /**
     * Sample code: WorkloadNetworks_GetVirtualMachine.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksGetVirtualMachine(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().getVirtualMachineWithResponse("group1", "cloud1", "vm1", Context.NONE);
    }
}
