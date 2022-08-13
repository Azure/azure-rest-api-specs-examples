import com.azure.core.util.Context;

/** Samples for Activity ListByModule. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listActivitiesByModule.json
     */
    /**
     * Sample code: List activities by a module.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listActivitiesByAModule(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.activities().listByModule("rg", "myAutomationAccount33", "OmsCompositeResources", Context.NONE);
    }
}
