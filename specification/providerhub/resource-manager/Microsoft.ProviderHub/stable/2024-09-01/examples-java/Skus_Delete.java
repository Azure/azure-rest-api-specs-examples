
/**
 * Samples for Skus Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Skus_Delete.json
     */
    /**
     * Sample code: Skus_Delete.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void skusDelete(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.skus().deleteWithResponse("Microsoft.Contoso", "testResourceType", "testSku",
            com.azure.core.util.Context.NONE);
    }
}
