
/**
 * Samples for CustomAssessmentAutomations GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/
     * CustomAssessmentAutomations/customAssessmentAutomationGet_example.json
     */
    /**
     * Sample code: Get a Custom Assessment Automation.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getACustomAssessmentAutomation(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.customAssessmentAutomations().getByResourceGroupWithResponse("TestResourceGroup",
            "MyCustomAssessmentAutomation", com.azure.core.util.Context.NONE);
    }
}
