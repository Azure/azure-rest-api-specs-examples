import com.azure.resourcemanager.devtestlabs.models.ResizeLabVirtualMachineProperties;

/** Samples for VirtualMachines Resize. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachines_Resize.json
     */
    /**
     * Sample code: VirtualMachines_Resize.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void virtualMachinesResize(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .virtualMachines()
            .resize(
                "resourceGroupName",
                "{labName}",
                "{vmName}",
                new ResizeLabVirtualMachineProperties().withSize("Standard_A4_v2"),
                com.azure.core.util.Context.NONE);
    }
}
