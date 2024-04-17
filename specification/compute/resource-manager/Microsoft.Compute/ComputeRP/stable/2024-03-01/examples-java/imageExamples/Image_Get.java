
/**
 * Samples for Images GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/imageExamples/
     * Image_Get.json
     */
    /**
     * Sample code: Get information about a virtual machine image.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getInformationAboutAVirtualMachineImage(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getImages().getByResourceGroupWithResponse("myResourceGroup",
            "myImage", null, com.azure.core.util.Context.NONE);
    }
}
