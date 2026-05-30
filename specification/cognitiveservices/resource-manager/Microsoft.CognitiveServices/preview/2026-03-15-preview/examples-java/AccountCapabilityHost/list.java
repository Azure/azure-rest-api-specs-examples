
/**
 * Samples for AccountCapabilityHosts List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/AccountCapabilityHost/list.json
     */
    /**
     * Sample code: List Account CapabilityHosts.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listAccountCapabilityHosts(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accountCapabilityHosts().list("test-rg", "account-1", com.azure.core.util.Context.NONE);
    }
}
