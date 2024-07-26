
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetReimageParameters;
import java.util.Arrays;

/**
 * Samples for VirtualMachineScaleSets Reimage.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSet_Reimage_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_Reimage_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualMachineScaleSetReimageMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSets().reimage("rgcompute",
            "aaaaaaaaaaaaaaaaaaaaaaaaaa", new VirtualMachineScaleSetReimageParameters().withTempDisk(true)
                .withForceUpdateOSDiskForEphemeral(true).withInstanceIds(Arrays.asList("aaaaaaaaaa")),
            com.azure.core.util.Context.NONE);
    }
}
