/** Samples for Skus DeleteNestedResourceTypeFirst. */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_DeleteNestedResourceTypeFirst.json
     */
    /**
     * Sample code: Skus_DeleteNestedResourceTypeFirst.
     *
     * @param manager Entry point to ProviderHubManager.
     */
    public static void skusDeleteNestedResourceTypeFirst(
        com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager
            .skus()
            .deleteNestedResourceTypeFirstWithResponse(
                "Microsoft.Contoso",
                "testResourceType",
                "nestedResourceTypeFirst",
                "testSku",
                com.azure.core.util.Context.NONE);
    }
}
