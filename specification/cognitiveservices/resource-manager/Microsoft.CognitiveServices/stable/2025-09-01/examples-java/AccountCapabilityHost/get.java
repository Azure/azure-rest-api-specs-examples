
/**
 * Samples for AccountCapabilityHosts Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/
     * AccountCapabilityHost/get.json
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
