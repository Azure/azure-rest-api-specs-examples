
import com.azure.resourcemanager.compute.models.VirtualMachineCaptureParameters;

/**
 * Samples for VirtualMachines Capture.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/
     * virtualMachineExamples/VirtualMachine_Capture_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachine_Capture_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineCaptureMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachines()
            .capture("rgcompute", "aaaaaaaaaaaaaaaaaaaa", new VirtualMachineCaptureParameters()
                .withVhdPrefix("aaaaaaaaa").withDestinationContainerName("aaaaaaa").withOverwriteVhds(true),
                com.azure.core.util.Context.NONE);
    }
}
