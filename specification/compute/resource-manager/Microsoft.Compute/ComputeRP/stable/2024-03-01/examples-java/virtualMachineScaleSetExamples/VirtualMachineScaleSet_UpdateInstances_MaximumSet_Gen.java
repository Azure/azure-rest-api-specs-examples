
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetVMInstanceRequiredIDs;
import java.util.Arrays;

/**
 * Samples for VirtualMachineScaleSets UpdateInstances.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSet_UpdateInstances_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_UpdateInstances_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualMachineScaleSetUpdateInstancesMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSets().updateInstances("rgcompute",
            "aaaaaaaaaaaaaaaaaaaaaaaa", new VirtualMachineScaleSetVMInstanceRequiredIDs()
                .withInstanceIds(Arrays.asList("aaaaaaaaaaaaaaaaaaaaaaaaa")),
            com.azure.core.util.Context.NONE);
    }
}
