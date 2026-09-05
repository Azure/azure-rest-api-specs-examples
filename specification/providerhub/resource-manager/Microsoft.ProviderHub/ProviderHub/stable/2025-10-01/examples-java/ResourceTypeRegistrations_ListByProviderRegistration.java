
/**
 * Samples for ResourceTypeRegistrations ListByProviderRegistration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/ResourceTypeRegistrations_ListByProviderRegistration.json
     */
    /**
     * Sample code: ResourceTypeRegistrations_ListByProviderRegistration.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void resourceTypeRegistrationsListByProviderRegistration(
        com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.resourceTypeRegistrations().listByProviderRegistration("Microsoft.Contoso",
            com.azure.core.util.Context.NONE);
    }
}
