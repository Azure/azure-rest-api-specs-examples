/** Samples for Skus DeleteNestedResourceTypeSecond. */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_DeleteNestedResourceTypeSecond.json
     */
    /**
     * Sample code: Skus_DeleteNestedResourceTypeSecond.
     *
     * @param manager Entry point to ProviderHubManager.
     */
    public static void skusDeleteNestedResourceTypeSecond(
        com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager
            .skus()
            .deleteNestedResourceTypeSecondWithResponse(
                "Microsoft.Contoso",
                "testResourceType",
                "nestedResourceTypeFirst",
                "nestedResourceTypeSecond",
                "testSku",
                com.azure.core.util.Context.NONE);
    }
}
