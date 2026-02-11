
/**
 * Samples for ApiKeys Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/ApiKeys_Delete.json
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
