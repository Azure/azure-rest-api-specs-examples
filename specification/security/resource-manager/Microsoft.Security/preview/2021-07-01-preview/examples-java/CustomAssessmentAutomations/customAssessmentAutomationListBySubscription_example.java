/** Samples for CustomAssessmentAutomations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/CustomAssessmentAutomations/customAssessmentAutomationListBySubscription_example.json
     */
    /**
     * Sample code: List Custom Assessment Automations in a subscription.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void listCustomAssessmentAutomationsInASubscription(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.customAssessmentAutomations().list(com.azure.core.util.Context.NONE);
    }
}
