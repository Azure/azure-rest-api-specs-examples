
/**
 * Samples for WebCategories Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/AzureWebCategoryGet.json
     */
    /**
     * Sample code: Get Azure Web Category by name.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getAzureWebCategoryByName(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getWebCategories().getWithResponse("Arts", null, com.azure.core.util.Context.NONE);
    }
}
