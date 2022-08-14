import com.azure.core.util.Context;

/** Samples for Schedule ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listSchedulesByAutomationAccount_First100.json
     */
    /**
     * Sample code: List schedules by automation account, first 100.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listSchedulesByAutomationAccountFirst100(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.schedules().listByAutomationAccount("rg", "myAutomationAccount33", Context.NONE);
    }
}
