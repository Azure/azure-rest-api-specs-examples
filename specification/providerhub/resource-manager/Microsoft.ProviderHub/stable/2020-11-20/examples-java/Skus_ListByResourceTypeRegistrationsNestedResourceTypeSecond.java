/** Samples for Skus ListByResourceTypeRegistrationsNestedResourceTypeSecond. */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_ListByResourceTypeRegistrationsNestedResourceTypeSecond.json
     */
    /**
     * Sample code: Skus_ListByResourceTypeRegistrationsNestedResourceTypeSecond.
     *
     * @param manager Entry point to ProviderHubManager.
     */
    public static void skusListByResourceTypeRegistrationsNestedResourceTypeSecond(
        com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager
            .skus()
            .listByResourceTypeRegistrationsNestedResourceTypeSecond(
                "Microsoft.Contoso",
                "testResourceType",
                "nestedResourceTypeFirst",
                "nestedResourceTypeSecond",
                com.azure.core.util.Context.NONE);
    }
}
