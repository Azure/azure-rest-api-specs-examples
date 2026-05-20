
/**
 * Samples for ProviderActions TriggerEvaluation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-27/TriggerEvaluation.json
     */
    /**
     * Sample code: TriggerEvaluation.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void
        triggerEvaluation(com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.providerActions().triggerEvaluation(null, com.azure.core.util.Context.NONE);
    }
}
