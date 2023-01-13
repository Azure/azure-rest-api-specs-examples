/** Samples for Products ListDetails. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestack/resource-manager/Microsoft.AzureStack/stable/2022-06-01/examples/Product/Post.json
     */
    /**
     * Sample code: Returns the extended properties of a product.
     *
     * @param manager Entry point to AzureStackManager.
     */
    public static void returnsTheExtendedPropertiesOfAProduct(
        com.azure.resourcemanager.azurestack.AzureStackManager manager) {
        manager
            .products()
            .listDetailsWithResponse(
                "azurestack",
                "testregistration",
                "Microsoft.OSTCExtensions.VMAccessForLinux.1.4.7.1",
                com.azure.core.util.Context.NONE);
    }
}
