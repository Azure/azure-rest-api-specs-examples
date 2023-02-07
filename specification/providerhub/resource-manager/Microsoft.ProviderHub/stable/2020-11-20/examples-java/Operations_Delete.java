/** Samples for Operations Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Operations_Delete.json
     */
    /**
     * Sample code: Operations_Delete.
     *
     * @param manager Entry point to ProviderHubManager.
     */
    public static void operationsDelete(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.operations().deleteWithResponse("Microsoft.Contoso", com.azure.core.util.Context.NONE);
    }
}
