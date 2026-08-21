
import com.azure.resourcemanager.compute.models.MigrateVMAvailabilityZoneInput;
import java.util.Arrays;

/**
 * Samples for VirtualMachineScaleSets MigrateVMAvailabilityZone.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-04-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_MigrateVMAvailabilityZone.json
     */
    /**
     * Sample code: VirtualMachineScaleSet Migrate VM Availability Zone.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetMigrateVMAvailabilityZone(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().migrateVMAvailabilityZone("myResourceGroup", "{vmss-name}",
            new MigrateVMAvailabilityZoneInput().withInstanceIds(Arrays.asList("0", "1", "2")).withTargetZone("2"),
            com.azure.core.util.Context.NONE);
    }
}
