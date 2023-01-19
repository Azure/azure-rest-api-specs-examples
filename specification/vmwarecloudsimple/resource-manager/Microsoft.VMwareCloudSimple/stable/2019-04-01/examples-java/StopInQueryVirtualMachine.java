import com.azure.resourcemanager.vmwarecloudsimple.models.StopMode;

/** Samples for VirtualMachines Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/StopInQueryVirtualMachine.json
     */
    /**
     * Sample code: StopInQueryVirtualMachine.
     *
     * @param manager Entry point to VMwareCloudSimpleManager.
     */
    public static void stopInQueryVirtualMachine(
        com.azure.resourcemanager.vmwarecloudsimple.VMwareCloudSimpleManager manager) {
        manager
            .virtualMachines()
            .stop(
                "myResourceGroup",
                "https://management.azure.com/",
                "myVirtualMachine",
                StopMode.SUSPEND,
                null,
                com.azure.core.util.Context.NONE);
    }
}
