import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetVMReimageParameters;

/** Samples for VirtualMachineScaleSetVMs Reimage. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSetVMs_Reimage_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVMs_Reimage_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetVMsReimageMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetVMs()
            .reimage(
                "rgcompute",
                "aaaaaaaaaaaaaaaa",
                "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
                new VirtualMachineScaleSetVMReimageParameters().withTempDisk(true),
                Context.NONE);
    }
}
