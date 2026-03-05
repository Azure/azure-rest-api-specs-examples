
/**
 * Samples for Skus DeleteNestedResourceTypeFirst.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Skus_DeleteNestedResourceTypeFirst.json
     */
    /**
     * Sample code: Skus_DeleteNestedResourceTypeFirst.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void
        skusDeleteNestedResourceTypeFirst(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.skus().deleteNestedResourceTypeFirstWithResponse("Microsoft.Contoso", "testResourceType",
            "nestedResourceTypeFirst", "testSku", com.azure.core.util.Context.NONE);
    }
}
