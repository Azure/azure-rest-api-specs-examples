/** Samples for Skus ListByResourceTypeRegistrationsNestedResourceTypeFirst. */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_ListByResourceTypeRegistrationsNestedResourceTypeFirst.json
     */
    /**
     * Sample code: Skus_ListByResourceTypeRegistrationsNestedResourceTypeFirst.
     *
     * @param manager Entry point to ProviderHubManager.
     */
    public static void skusListByResourceTypeRegistrationsNestedResourceTypeFirst(
        com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager
            .skus()
            .listByResourceTypeRegistrationsNestedResourceTypeFirst(
                "Microsoft.Contoso", "testResourceType", "nestedResourceTypeFirst", com.azure.core.util.Context.NONE);
    }
}
