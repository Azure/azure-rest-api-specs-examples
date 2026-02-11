
/**
 * Samples for ApiKeys List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/ApiKeys_List.json
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
