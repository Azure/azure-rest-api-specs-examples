import com.azure.core.util.Context;

/** Samples for Variable ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listVariables_Next100.json
     */
    /**
     * Sample code: List variables, Next 100.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listVariablesNext100(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.variables().listByAutomationAccount("rg", "sampleAccount9", Context.NONE);
    }
}
