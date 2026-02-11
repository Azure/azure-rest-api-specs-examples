
/**
 * Samples for ApiKeys Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/ApiKeys_Get.json
     */
    /**
     * Sample code: ApiKeys_Get.
     * 
     * @param manager Entry point to NginxManager.
     */
    public static void apiKeysGet(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.apiKeys().getWithResponse("myResourceGroup", "myDeployment", "myApiKey",
            com.azure.core.util.Context.NONE);
    }
}
