import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.VirtualMachineCaptureParameters;

/** Samples for VirtualMachines Capture. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineExamples/VirtualMachines_Capture_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachines_Capture_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachinesCaptureMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachines()
            .capture(
                "rgcompute",
                "aaaaaaaaaaaaa",
                new VirtualMachineCaptureParameters()
                    .withVhdPrefix("aaaaaaaaa")
                    .withDestinationContainerName("aaaaaaa")
                    .withOverwriteVhds(true),
                Context.NONE);
    }
}
