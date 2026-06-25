
import com.azure.resourcemanager.compute.models.VirtualMachineCaptureParameters;

/**
 * Samples for VirtualMachines Capture.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_Capture_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachine_Capture_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineCaptureMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().capture("rgcompute", "aaaaaaaaaaaaa",
            new VirtualMachineCaptureParameters().withVhdPrefix("aaaaaaaaa").withDestinationContainerName("aaaaaaa")
                .withOverwriteVhds(true),
            com.azure.core.util.Context.NONE);
    }
}
