/** Samples for Products Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestack/resource-manager/Microsoft.AzureStack/stable/2022-06-01/examples/Product/Get.json
     */
    /**
     * Sample code: Returns the specified product.
     *
     * @param manager Entry point to AzureStackManager.
     */
    public static void returnsTheSpecifiedProduct(com.azure.resourcemanager.azurestack.AzureStackManager manager) {
        manager
            .products()
            .getWithResponse(
                "azurestack",
                "testregistration",
                "Microsoft.OSTCExtensions.VMAccessForLinux.1.4.7.1",
                com.azure.core.util.Context.NONE);
    }
}
