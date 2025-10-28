
/**
 * Samples for AccountCapabilityHosts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/
     * AccountCapabilityHost/delete.json
     */
    /**
     * Sample code: Delete Account CapabilityHost.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        deleteAccountCapabilityHost(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accountCapabilityHosts().delete("test-rg", "account-1", "capabilityHostName",
            com.azure.core.util.Context.NONE);
    }
}
