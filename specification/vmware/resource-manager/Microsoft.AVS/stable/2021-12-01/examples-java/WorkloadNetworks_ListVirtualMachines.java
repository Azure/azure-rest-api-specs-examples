import com.azure.core.util.Context;

/** Samples for WorkloadNetworks ListVirtualMachines. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_ListVirtualMachines.json
     */
    /**
     * Sample code: WorkloadNetworks_ListVirtualMachines.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksListVirtualMachines(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().listVirtualMachines("group1", "cloud1", Context.NONE);
    }
}
