import com.azure.core.util.Context;

/** Samples for Module Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/deleteModule.json
     */
    /**
     * Sample code: Delete a module.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void deleteAModule(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.modules().deleteWithResponse("rg", "myAutomationAccount33", "OmsCompositeResources", Context.NONE);
    }
}
