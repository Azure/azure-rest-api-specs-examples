
/**
 * Samples for Operations ListByProviderRegistration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/Operations_ListByProviderRegistration.json
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
