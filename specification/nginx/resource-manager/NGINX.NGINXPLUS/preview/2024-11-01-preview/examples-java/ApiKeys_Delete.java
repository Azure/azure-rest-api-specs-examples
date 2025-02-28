
/**
 * Samples for ApiKeys Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/nginx/resource-manager/NGINX.NGINXPLUS/preview/2024-11-01-preview/examples/ApiKeys_Delete.json
     */
    /**
     * Sample code: ApiKeys_Delete.
     * 
     * @param manager Entry point to NginxManager.
     */
    public static void apiKeysDelete(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.apiKeys().deleteWithResponse("myResourceGroup", "myDeployment", "myApiKey",
            com.azure.core.util.Context.NONE);
    }
}
