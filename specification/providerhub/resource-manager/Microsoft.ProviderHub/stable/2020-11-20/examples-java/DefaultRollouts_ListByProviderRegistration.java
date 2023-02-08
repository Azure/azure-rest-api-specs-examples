/** Samples for DefaultRollouts ListByProviderRegistration. */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/DefaultRollouts_ListByProviderRegistration.json
     */
    /**
     * Sample code: DefaultRollouts_ListByProviderRegistration.
     *
     * @param manager Entry point to ProviderHubManager.
     */
    public static void defaultRolloutsListByProviderRegistration(
        com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.defaultRollouts().listByProviderRegistration("Microsoft.Contoso", com.azure.core.util.Context.NONE);
    }
}
