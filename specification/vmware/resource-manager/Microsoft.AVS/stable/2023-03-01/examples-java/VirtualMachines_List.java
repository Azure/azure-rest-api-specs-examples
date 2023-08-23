/** Samples for VirtualMachines List. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/VirtualMachines_List.json
     */
    /**
     * Sample code: ListClusterVirtualMachines.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void listClusterVirtualMachines(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.virtualMachines().list("group1", "cloud1", "cluster1", com.azure.core.util.Context.NONE);
    }
}
