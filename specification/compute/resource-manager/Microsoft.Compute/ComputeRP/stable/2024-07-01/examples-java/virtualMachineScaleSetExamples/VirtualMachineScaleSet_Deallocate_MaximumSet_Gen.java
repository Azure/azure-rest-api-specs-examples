
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetVMInstanceIDs;
import java.util.Arrays;

/**
 * Samples for VirtualMachineScaleSets Deallocate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSet_Deallocate_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_Deallocate_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualMachineScaleSetDeallocateMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSets().deallocate("rgcompute",
            "aaaaaaaaaaaaaaaaaaaaaaaaaaaa", true,
            new VirtualMachineScaleSetVMInstanceIDs().withInstanceIds(Arrays.asList("aaaaaaaaaaaaaaaaa")),
            com.azure.core.util.Context.NONE);
    }
}
