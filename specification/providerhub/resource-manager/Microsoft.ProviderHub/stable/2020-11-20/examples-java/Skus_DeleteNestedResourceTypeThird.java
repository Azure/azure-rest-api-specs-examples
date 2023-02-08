/** Samples for Skus DeleteNestedResourceTypeThird. */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_DeleteNestedResourceTypeThird.json
     */
    /**
     * Sample code: Skus_DeleteNestedResourceTypeThird.
     *
     * @param manager Entry point to ProviderHubManager.
     */
    public static void skusDeleteNestedResourceTypeThird(
        com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager
            .skus()
            .deleteNestedResourceTypeThirdWithResponse(
                "Microsoft.Contoso",
                "testResourceType",
                "nestedResourceTypeFirst",
                "nestedResourceTypeSecond",
                "nestedResourceTypeThird",
                "testSku",
                com.azure.core.util.Context.NONE);
    }
}
