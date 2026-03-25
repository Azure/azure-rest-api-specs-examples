
/**
 * Samples for ApiKeys CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/ApiKeys_CreateOrUpdate.json
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
