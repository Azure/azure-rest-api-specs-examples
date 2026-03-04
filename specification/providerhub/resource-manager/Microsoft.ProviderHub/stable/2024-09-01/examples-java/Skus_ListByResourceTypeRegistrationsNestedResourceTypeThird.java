
/**
 * Samples for Skus ListByResourceTypeRegistrationsNestedResourceTypeThird.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Skus_ListByResourceTypeRegistrationsNestedResourceTypeThird.json
     */
    /**
     * Sample code: Skus_ListByResourceTypeRegistrationsNestedResourceTypeThird.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void skusListByResourceTypeRegistrationsNestedResourceTypeThird(
        com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.skus().listByResourceTypeRegistrationsNestedResourceTypeThird("Microsoft.Contoso", "testResourceType",
            "nestedResourceTypeFirst", "nestedResourceTypeSecond", "nestedResourceTypeThird",
            com.azure.core.util.Context.NONE);
    }
}
