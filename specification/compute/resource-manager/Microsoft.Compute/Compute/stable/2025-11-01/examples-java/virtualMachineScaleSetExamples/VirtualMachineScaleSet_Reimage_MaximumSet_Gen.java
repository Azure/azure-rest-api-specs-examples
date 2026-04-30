
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetReimageParameters;
import java.util.Arrays;

/**
 * Samples for VirtualMachineScaleSets Reimage.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Reimage_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_Reimage_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetReimageMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().reimage(
            "rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaaaa", new VirtualMachineScaleSetReimageParameters().withTempDisk(true)
                .withForceUpdateOSDiskForEphemeral(true).withInstanceIds(Arrays.asList("aaaaaaaaaa")),
            com.azure.core.util.Context.NONE);
    }
}
