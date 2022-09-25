import com.azure.core.util.Context;

/** Samples for VirtualMachineImages Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineImageExamples/VirtualMachineImages_Get_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImages_Get_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineImagesGetMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineImages()
            .getWithResponse(
                "aaaaaaaaaaaa", "aaaaaaaaaaa", "aa", "aaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa", Context.NONE);
    }
}
