/** Samples for VirtualMachines ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/ListRGVirtualMachines.json
     */
    /**
     * Sample code: ListRGVirtualMachines.
     *
     * @param manager Entry point to VMwareCloudSimpleManager.
     */
    public static void listRGVirtualMachines(
        com.azure.resourcemanager.vmwarecloudsimple.VMwareCloudSimpleManager manager) {
        manager
            .virtualMachines()
            .listByResourceGroup("myResourceGroup", null, null, null, com.azure.core.util.Context.NONE);
    }
}
