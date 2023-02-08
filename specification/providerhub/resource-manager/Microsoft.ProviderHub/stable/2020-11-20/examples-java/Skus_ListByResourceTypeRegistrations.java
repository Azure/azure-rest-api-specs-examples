/** Samples for Skus ListByResourceTypeRegistrations. */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_ListByResourceTypeRegistrations.json
     */
    /**
     * Sample code: Skus_ListByResourceTypeRegistrations.
     *
     * @param manager Entry point to ProviderHubManager.
     */
    public static void skusListByResourceTypeRegistrations(
        com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager
            .skus()
            .listByResourceTypeRegistrations("Microsoft.Contoso", "testResourceType", com.azure.core.util.Context.NONE);
    }
}
