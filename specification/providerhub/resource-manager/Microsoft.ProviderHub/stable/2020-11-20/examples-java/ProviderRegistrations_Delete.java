/** Samples for ProviderRegistrations Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/ProviderRegistrations_Delete.json
     */
    /**
     * Sample code: ProviderRegistrations_Delete.
     *
     * @param manager Entry point to ProviderHubManager.
     */
    public static void providerRegistrationsDelete(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.providerRegistrations().deleteWithResponse("Microsoft.Contoso", com.azure.core.util.Context.NONE);
    }
}
