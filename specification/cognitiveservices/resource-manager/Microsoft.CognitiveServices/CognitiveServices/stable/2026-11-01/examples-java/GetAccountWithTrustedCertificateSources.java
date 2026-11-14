
/**
 * Samples for Accounts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/GetAccountWithTrustedCertificateSources.json
     */
    /**
     * Sample code: Get an account with private CA trust.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getAnAccountWithPrivateCATrust(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accounts().getByResourceGroupWithResponse("myResourceGroup", "myAccount",
            com.azure.core.util.Context.NONE);
    }
}
