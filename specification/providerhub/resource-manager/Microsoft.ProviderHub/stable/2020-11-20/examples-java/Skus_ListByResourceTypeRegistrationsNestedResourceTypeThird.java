/** Samples for Skus ListByResourceTypeRegistrationsNestedResourceTypeThird. */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_ListByResourceTypeRegistrationsNestedResourceTypeThird.json
     */
    /**
     * Sample code: Skus_ListByResourceTypeRegistrationsNestedResourceTypeThird.
     *
     * @param manager Entry point to ProviderHubManager.
     */
    public static void skusListByResourceTypeRegistrationsNestedResourceTypeThird(
        com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager
            .skus()
            .listByResourceTypeRegistrationsNestedResourceTypeThird(
                "Microsoft.Contoso",
                "testResourceType",
                "nestedResourceTypeFirst",
                "nestedResourceTypeSecond",
                "nestedResourceTypeThird",
                com.azure.core.util.Context.NONE);
    }
}
