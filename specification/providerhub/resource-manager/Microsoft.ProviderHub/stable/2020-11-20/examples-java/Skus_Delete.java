/** Samples for Skus Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_Delete.json
     */
    /**
     * Sample code: Skus_Delete.
     *
     * @param manager Entry point to ProviderHubManager.
     */
    public static void skusDelete(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager
            .skus()
            .deleteWithResponse("Microsoft.Contoso", "testResourceType", "testSku", com.azure.core.util.Context.NONE);
    }
}
