import com.azure.core.util.Context;

/** Samples for Images GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/GetInformationAboutAnImage.json
     */
    /**
     * Sample code: Get information about a virtual machine image.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getInformationAboutAVirtualMachineImage(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getImages()
            .getByResourceGroupWithResponse("myResourceGroup", "myImage", null, Context.NONE);
    }
}
