
/**
 * Samples for Skus ListByResourceTypeRegistrations.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Skus_ListByResourceTypeRegistrations.json
     */
    /**
     * Sample code: Skus_ListByResourceTypeRegistrations.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void
        skusListByResourceTypeRegistrations(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.skus().listByResourceTypeRegistrations("Microsoft.Contoso", "testResourceType",
            com.azure.core.util.Context.NONE);
    }
}
