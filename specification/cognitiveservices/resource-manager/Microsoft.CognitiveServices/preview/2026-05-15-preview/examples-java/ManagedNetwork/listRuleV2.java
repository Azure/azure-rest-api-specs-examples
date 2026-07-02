
/**
 * Samples for OutboundRule List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/ManagedNetwork/listRuleV2.json
     */
    /**
     * Sample code: List OutboundRules.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listOutboundRules(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.outboundRules().list("test-rg", "cognitive-account-name", "default", com.azure.core.util.Context.NONE);
    }
}
