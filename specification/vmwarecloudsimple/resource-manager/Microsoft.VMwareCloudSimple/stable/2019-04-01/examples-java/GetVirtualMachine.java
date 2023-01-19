/** Samples for VirtualMachines GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/GetVirtualMachine.json
     */
    /**
     * Sample code: GetVirtualMachine.
     *
     * @param manager Entry point to VMwareCloudSimpleManager.
     */
    public static void getVirtualMachine(com.azure.resourcemanager.vmwarecloudsimple.VMwareCloudSimpleManager manager) {
        manager
            .virtualMachines()
            .getByResourceGroupWithResponse("myResourceGroup", "myVirtualMachine", com.azure.core.util.Context.NONE);
    }
}
