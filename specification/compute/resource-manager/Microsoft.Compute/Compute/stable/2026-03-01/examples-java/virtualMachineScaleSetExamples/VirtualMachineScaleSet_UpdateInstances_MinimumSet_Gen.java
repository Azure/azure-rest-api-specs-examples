
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetVMInstanceRequiredIDs;
import java.util.Arrays;

/**
 * Samples for VirtualMachineScaleSets UpdateInstances.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_UpdateInstances_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_UpdateInstances_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetUpdateInstancesMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().updateInstances("rgcompute",
            "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa", new VirtualMachineScaleSetVMInstanceRequiredIDs()
                .withInstanceIds(Arrays.asList("aaaaaaaaaaaaaaaaaaaaaaaaa")),
            com.azure.core.util.Context.NONE);
    }
}
