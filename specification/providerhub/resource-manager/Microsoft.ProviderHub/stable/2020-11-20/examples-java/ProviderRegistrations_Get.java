/** Samples for ProviderRegistrations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/ProviderRegistrations_Get.json
     */
    /**
     * Sample code: ProviderRegistrations_Get.
     *
     * @param manager Entry point to ProviderHubManager.
     */
    public static void providerRegistrationsGet(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.providerRegistrations().getWithResponse("Microsoft.Contoso", com.azure.core.util.Context.NONE);
    }
}
