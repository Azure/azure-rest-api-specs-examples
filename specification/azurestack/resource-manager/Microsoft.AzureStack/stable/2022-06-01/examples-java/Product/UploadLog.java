/** Samples for Products UploadLog. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestack/resource-manager/Microsoft.AzureStack/stable/2022-06-01/examples/Product/UploadLog.json
     */
    /**
     * Sample code: Returns the specified product.
     *
     * @param manager Entry point to AzureStackManager.
     */
    public static void returnsTheSpecifiedProduct(com.azure.resourcemanager.azurestack.AzureStackManager manager) {
        manager
            .products()
            .uploadLogWithResponse(
                "azurestack",
                "testregistration",
                "Microsoft.OSTCExtensions.VMAccessForLinux.1.4.7.1",
                null,
                com.azure.core.util.Context.NONE);
    }
}
