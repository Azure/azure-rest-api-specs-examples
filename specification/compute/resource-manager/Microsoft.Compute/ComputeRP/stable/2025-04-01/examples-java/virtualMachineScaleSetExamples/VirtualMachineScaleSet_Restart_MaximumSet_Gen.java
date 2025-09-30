
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetVMInstanceIDs;
import java.util.Arrays;

/**
 * Samples for VirtualMachineScaleSets Restart.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSet_Restart_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_Restart_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualMachineScaleSetRestartMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSets().restart("rgcompute",
            "aaaaaaaaaaaaaaaaaaa",
            new VirtualMachineScaleSetVMInstanceIDs().withInstanceIds(Arrays.asList("aaaaaaaaaaaaaaaaa")),
            com.azure.core.util.Context.NONE);
    }
}
