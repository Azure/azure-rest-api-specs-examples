/** Samples for ProviderRegistrations GenerateOperations. */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/ProviderRegistrations_GenerateOperations.json
     */
    /**
     * Sample code: ProviderRegistrations_GenerateOperations.
     *
     * @param manager Entry point to ProviderHubManager.
     */
    public static void providerRegistrationsGenerateOperations(
        com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager
            .providerRegistrations()
            .generateOperationsWithResponse("Microsoft.Contoso", com.azure.core.util.Context.NONE);
    }
}
