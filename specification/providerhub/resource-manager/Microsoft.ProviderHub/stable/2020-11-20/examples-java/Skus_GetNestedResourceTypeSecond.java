/** Samples for Skus GetNestedResourceTypeSecond. */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_GetNestedResourceTypeSecond.json
     */
    /**
     * Sample code: Skus_GetNestedResourceTypeSecond.
     *
     * @param manager Entry point to ProviderHubManager.
     */
    public static void skusGetNestedResourceTypeSecond(
        com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager
            .skus()
            .getNestedResourceTypeSecondWithResponse(
                "Microsoft.Contoso",
                "testResourceType",
                "nestedResourceTypeFirst",
                "nestedResourceTypeSecond",
                "testSku",
                com.azure.core.util.Context.NONE);
    }
}
