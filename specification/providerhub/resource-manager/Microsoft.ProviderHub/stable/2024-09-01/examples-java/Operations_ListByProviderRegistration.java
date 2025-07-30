
/**
 * Samples for Operations ListByProviderRegistration.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/
     * Operations_ListByProviderRegistration.json
     */
    /**
     * Sample code: Operations_ListByProviderRegistration.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void
        operationsListByProviderRegistration(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.operations().listByProviderRegistrationWithResponse("Microsoft.Contoso",
            com.azure.core.util.Context.NONE);
    }
}
