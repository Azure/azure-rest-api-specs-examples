import com.azure.core.util.Context;

/** Samples for Activity Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getActivityInAModule.json
     */
    /**
     * Sample code: Get Activity in a module.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getActivityInAModule(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .activities()
            .getWithResponse(
                "rg", "myAutomationAccount33", "OmsCompositeResources", "Add-AzureRmAccount", Context.NONE);
    }
}
