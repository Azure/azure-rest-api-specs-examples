
/**
 * Samples for CustomAssessmentAutomations ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/
     * CustomAssessmentAutomations/customAssessmentAutomationListByResourceGroup_example.json
     */
    /**
     * Sample code: List Custom Assessment Automations in a subscription and a resource group.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void listCustomAssessmentAutomationsInASubscriptionAndAResourceGroup(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.customAssessmentAutomations().listByResourceGroup("TestResourceGroup",
            com.azure.core.util.Context.NONE);
    }
}
