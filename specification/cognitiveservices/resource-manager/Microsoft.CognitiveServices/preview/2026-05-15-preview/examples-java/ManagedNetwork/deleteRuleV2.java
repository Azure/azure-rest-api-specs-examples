
/**
 * Samples for OutboundRule Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/ManagedNetwork/deleteRuleV2.json
     */
    /**
     * Sample code: Delete OutboundRule.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        deleteOutboundRule(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.outboundRules().delete("test-rg", "cognitive-account-name", "default", "rule-name",
            com.azure.core.util.Context.NONE);
    }
}
