
/**
 * Samples for Skus List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/Skus_List.json
     */
    /**
     * Sample code: Skus_List.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void skusList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getSkus().list(com.azure.core.util.Context.NONE);
    }
}
