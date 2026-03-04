
/**
 * Samples for Skus GetNestedResourceTypeThird.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Skus_GetNestedResourceTypeThird.json
     */
    /**
     * Sample code: Skus_GetNestedResourceTypeThird.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void
        skusGetNestedResourceTypeThird(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.skus().getNestedResourceTypeThirdWithResponse("Microsoft.Contoso", "testResourceType",
            "nestedResourceTypeFirst", "nestedResourceTypeSecond", "nestedResourceTypeThird", "testSku",
            com.azure.core.util.Context.NONE);
    }
}
