
/**
 * Samples for OutboundRule Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/ManagedNetwork/getRuleV2.json
     */
    /**
     * Sample code: Get OutboundRule.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void getOutboundRule(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.outboundRules().getWithResponse("test-rg", "cognitive-account-name", "default", "name_of_the_fqdn_rule",
            com.azure.core.util.Context.NONE);
    }
}
