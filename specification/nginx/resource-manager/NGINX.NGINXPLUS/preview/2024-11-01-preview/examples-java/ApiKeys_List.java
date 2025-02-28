
/**
 * Samples for ApiKeys List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/nginx/resource-manager/NGINX.NGINXPLUS/preview/2024-11-01-preview/examples/ApiKeys_List.json
     */
    /**
     * Sample code: ApiKeys_List.
     * 
     * @param manager Entry point to NginxManager.
     */
    public static void apiKeysList(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.apiKeys().list("myResourceGroup", "myDeployment", com.azure.core.util.Context.NONE);
    }
}
