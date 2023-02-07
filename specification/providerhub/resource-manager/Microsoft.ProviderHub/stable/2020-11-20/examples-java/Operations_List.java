/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     *
     * @param manager Entry point to ProviderHubManager.
     */
    public static void operationsList(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
