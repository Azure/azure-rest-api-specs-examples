
/**
 * Samples for CustomRollouts ListByProviderRegistration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/CustomRollouts_ListByProviderRegistration.json
     */
    /**
     * Sample code: CustomRollouts_ListByProviderRegistration.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void
        customRolloutsListByProviderRegistration(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.customRollouts().listByProviderRegistration("Microsoft.Contoso", com.azure.core.util.Context.NONE);
    }
}
