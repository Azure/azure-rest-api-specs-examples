import com.azure.core.util.Context;

/** Samples for VirtualMachineExtensionImages Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineExtensionImageExamples/VirtualMachineExtensionImages_Get_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineExtensionImages_Get_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineExtensionImagesGetMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineExtensionImages()
            .getWithResponse("aaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaa", "aa", "aaa", Context.NONE);
    }
}
