/** Samples for VirtualMachines Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/VirtualMachines_Get.json
     */
    /**
     * Sample code: GetVirtualMachine.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void getVirtualMachine(com.azure.resourcemanager.avs.AvsManager manager) {
        manager
            .virtualMachines()
            .getWithResponse("group1", "cloud1", "cluster1", "vm-209", com.azure.core.util.Context.NONE);
    }
}
