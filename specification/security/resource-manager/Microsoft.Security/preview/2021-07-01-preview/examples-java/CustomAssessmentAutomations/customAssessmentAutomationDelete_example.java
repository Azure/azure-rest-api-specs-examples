
/**
 * Samples for CustomAssessmentAutomations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/
     * CustomAssessmentAutomations/customAssessmentAutomationDelete_example.json
     */
    /**
     * Sample code: Delete a Custom Assessment Automation.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void deleteACustomAssessmentAutomation(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.customAssessmentAutomations().deleteByResourceGroupWithResponse("TestResourceGroup",
            "MyCustomAssessmentAutomation", com.azure.core.util.Context.NONE);
    }
}
