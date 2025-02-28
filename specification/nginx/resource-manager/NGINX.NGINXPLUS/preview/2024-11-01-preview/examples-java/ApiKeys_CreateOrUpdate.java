
/**
 * Samples for ApiKeys CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/nginx/resource-manager/NGINX.NGINXPLUS/preview/2024-11-01-preview/examples/ApiKeys_CreateOrUpdate.
     * json
     */
    /**
     * Sample code: ApiKeys_CreateOrUpdate.
     * 
     * @param manager Entry point to NginxManager.
     */
    public static void apiKeysCreateOrUpdate(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.apiKeys().define("myApiKey").withExistingNginxDeployment("myResourceGroup", "myDeployment").create();
    }
}
