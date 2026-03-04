
/**
 * Samples for Skus Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Skus_Get.json
     */
    /**
     * Sample code: Skus_Get.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void skusGet(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.skus().getWithResponse("Microsoft.Contoso", "testResourceType", "testSku",
            com.azure.core.util.Context.NONE);
    }
}
