
/**
 * Samples for Module Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getModule.json
     */
    /**
     * Sample code: Get a module.
     * 
     * @param manager Entry point to AutomationManager.
     */
    public static void getAModule(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.modules().getWithResponse("rg", "myAutomationAccount33", "OmsCompositeResources",
            com.azure.core.util.Context.NONE);
    }
}
