
/**
 * Samples for Skus DeleteNestedResourceTypeSecond.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Skus_DeleteNestedResourceTypeSecond.json
     */
    /**
     * Sample code: Skus_DeleteNestedResourceTypeSecond.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void
        skusDeleteNestedResourceTypeSecond(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.skus().deleteNestedResourceTypeSecondWithResponse("Microsoft.Contoso", "testResourceType",
            "nestedResourceTypeFirst", "nestedResourceTypeSecond", "testSku", com.azure.core.util.Context.NONE);
    }
}
