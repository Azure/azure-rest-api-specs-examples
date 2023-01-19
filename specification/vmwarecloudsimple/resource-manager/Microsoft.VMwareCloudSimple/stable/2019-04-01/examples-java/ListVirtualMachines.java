/** Samples for VirtualMachines List. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/ListVirtualMachines.json
     */
    /**
     * Sample code: ListVirtualMachines.
     *
     * @param manager Entry point to VMwareCloudSimpleManager.
     */
    public static void listVirtualMachines(
        com.azure.resourcemanager.vmwarecloudsimple.VMwareCloudSimpleManager manager) {
        manager.virtualMachines().list(null, null, null, com.azure.core.util.Context.NONE);
    }
}
