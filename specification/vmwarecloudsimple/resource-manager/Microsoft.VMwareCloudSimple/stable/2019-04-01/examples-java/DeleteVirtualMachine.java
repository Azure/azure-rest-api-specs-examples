/** Samples for VirtualMachines Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/DeleteVirtualMachine.json
     */
    /**
     * Sample code: DeleteVirtualMachine.
     *
     * @param manager Entry point to VMwareCloudSimpleManager.
     */
    public static void deleteVirtualMachine(
        com.azure.resourcemanager.vmwarecloudsimple.VMwareCloudSimpleManager manager) {
        manager
            .virtualMachines()
            .delete(
                "myResourceGroup",
                "https://management.azure.com/",
                "myVirtualMachine",
                com.azure.core.util.Context.NONE);
    }
}
