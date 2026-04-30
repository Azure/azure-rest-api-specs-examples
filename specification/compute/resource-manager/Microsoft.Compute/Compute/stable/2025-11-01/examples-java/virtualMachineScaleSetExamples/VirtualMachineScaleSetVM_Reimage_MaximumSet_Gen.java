
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetVMReimageParameters;

/**
 * Samples for VirtualMachineScaleSetVMs Reimage.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Reimage_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVM_Reimage_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetVMReimageMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMs().reimage("rgcompute", "aaaaaaaaaaaaaaaa",
            "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
            new VirtualMachineScaleSetVMReimageParameters().withTempDisk(true).withForceUpdateOSDiskForEphemeral(true),
            com.azure.core.util.Context.NONE);
    }
}
