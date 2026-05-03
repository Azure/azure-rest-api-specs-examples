
/**
 * Samples for AccountCapabilityHosts Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-15-preview/AccountCapabilityHost/get.json
     */
    /**
     * Sample code: Get Account CapabilityHost.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getAccountCapabilityHost(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accountCapabilityHosts().getWithResponse("test-rg", "account-1", "capabilityHostName",
            com.azure.core.util.Context.NONE);
    }
}
