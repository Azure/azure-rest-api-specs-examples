
import com.azure.resourcemanager.compute.models.MigrateVMToVirtualMachineScaleSetInput;

/**
 * Samples for VirtualMachines MigrateToVMScaleSet.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_MigrateToVirtualMachineScaleSet.json
     */
    /**
     * Sample code: Migrate a Virtual Machine to Flexible Virtual Machine Scale Ser.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void migrateAVirtualMachineToFlexibleVirtualMachineScaleSer(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().migrateToVMScaleSet("myResourceGroup", "myVMName",
            new MigrateVMToVirtualMachineScaleSetInput().withTargetFaultDomain(0).withTargetVMSize("Standard_D1_v2"),
            com.azure.core.util.Context.NONE);
    }
}
