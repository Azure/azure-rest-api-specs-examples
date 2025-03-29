
import com.azure.resourcemanager.compute.models.MigrateVMToVirtualMachineScaleSetInput;

/**
 * Samples for VirtualMachines MigrateToVMScaleSet.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * virtualMachineExamples/VirtualMachine_MigrateToVirtualMachineScaleSet.json
     */
    /**
     * Sample code: Migrate a Virtual Machine to Flexible Virtual Machine Scale Ser.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        migrateAVirtualMachineToFlexibleVirtualMachineScaleSer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachines().migrateToVMScaleSet("myResourceGroup",
            "myVMName",
            new MigrateVMToVirtualMachineScaleSetInput().withTargetFaultDomain(0).withTargetVMSize("Standard_D1_v2"),
            com.azure.core.util.Context.NONE);
    }
}
