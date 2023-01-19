import com.azure.resourcemanager.vmwarecloudsimple.models.VirtualMachineStopMode;

/** Samples for VirtualMachines Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/StopInBodyVirtualMachine.json
     */
    /**
     * Sample code: StopInBodyVirtualMachine.
     *
     * @param manager Entry point to VMwareCloudSimpleManager.
     */
    public static void stopInBodyVirtualMachine(
        com.azure.resourcemanager.vmwarecloudsimple.VMwareCloudSimpleManager manager) {
        manager
            .virtualMachines()
            .stop(
                "myResourceGroup",
                "https://management.azure.com/",
                "myVirtualMachine",
                null,
                new VirtualMachineStopMode(),
                com.azure.core.util.Context.NONE);
    }
}
