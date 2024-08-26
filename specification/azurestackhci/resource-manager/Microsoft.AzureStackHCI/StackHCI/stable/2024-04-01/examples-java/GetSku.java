
/**
 * Samples for Skus Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/GetSku.
     * json
     */
    /**
     * Sample code: Get Sku.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void getSku(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.skus().getWithResponse("test-rg", "myCluster", "publisher1", "offer1", "sku1", null,
            com.azure.core.util.Context.NONE);
    }
}
