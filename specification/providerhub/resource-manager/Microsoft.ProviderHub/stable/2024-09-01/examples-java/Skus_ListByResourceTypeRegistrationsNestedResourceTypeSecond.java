
/**
 * Samples for Skus ListByResourceTypeRegistrationsNestedResourceTypeSecond.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Skus_ListByResourceTypeRegistrationsNestedResourceTypeSecond.json
     */
    /**
     * Sample code: Skus_ListByResourceTypeRegistrationsNestedResourceTypeSecond.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void skusListByResourceTypeRegistrationsNestedResourceTypeSecond(
        com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.skus().listByResourceTypeRegistrationsNestedResourceTypeSecond("Microsoft.Contoso", "testResourceType",
            "nestedResourceTypeFirst", "nestedResourceTypeSecond", com.azure.core.util.Context.NONE);
    }
}
