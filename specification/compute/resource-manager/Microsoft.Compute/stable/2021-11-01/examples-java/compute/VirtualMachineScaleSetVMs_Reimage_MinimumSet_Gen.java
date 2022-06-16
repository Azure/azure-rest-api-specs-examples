import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSetVMs Reimage. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSetVMs_Reimage_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVMs_Reimage_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetVMsReimageMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetVMs()
            .reimage("rgcompute", "aaaaaaa", "aaaaaaaaaaaaa", null, Context.NONE);
    }
}
