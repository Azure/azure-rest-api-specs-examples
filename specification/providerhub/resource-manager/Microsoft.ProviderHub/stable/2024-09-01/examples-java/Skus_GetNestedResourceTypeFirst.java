
/**
 * Samples for Skus GetNestedResourceTypeFirst.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Skus_GetNestedResourceTypeFirst.json
     */
    /**
     * Sample code: Skus_GetNestedResourceTypeFirst.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void
        skusGetNestedResourceTypeFirst(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.skus().getNestedResourceTypeFirstWithResponse("Microsoft.Contoso", "testResourceType",
            "nestedResourceTypeFirst", "testSku", com.azure.core.util.Context.NONE);
    }
}
