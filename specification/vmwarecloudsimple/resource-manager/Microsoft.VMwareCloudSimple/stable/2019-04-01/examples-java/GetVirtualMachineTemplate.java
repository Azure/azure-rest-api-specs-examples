/** Samples for VirtualMachineTemplates Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/GetVirtualMachineTemplate.json
     */
    /**
     * Sample code: GetVirtualMachineTemplate.
     *
     * @param manager Entry point to VMwareCloudSimpleManager.
     */
    public static void getVirtualMachineTemplate(
        com.azure.resourcemanager.vmwarecloudsimple.VMwareCloudSimpleManager manager) {
        manager
            .virtualMachineTemplates()
            .getWithResponse("westus2", "myPrivateCloud", "vm-34", com.azure.core.util.Context.NONE);
    }
}
