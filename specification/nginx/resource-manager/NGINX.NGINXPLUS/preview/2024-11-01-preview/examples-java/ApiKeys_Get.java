
/**
 * Samples for ApiKeys Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/nginx/resource-manager/NGINX.NGINXPLUS/preview/2024-11-01-preview/examples/ApiKeys_Get.json
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
