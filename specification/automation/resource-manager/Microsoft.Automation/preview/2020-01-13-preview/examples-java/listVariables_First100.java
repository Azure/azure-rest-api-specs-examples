import com.azure.core.util.Context;

/** Samples for Variable ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listVariables_First100.json
     */
    /**
     * Sample code: List variables, First 100.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listVariablesFirst100(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.variables().listByAutomationAccount("rg", "sampleAccount9", Context.NONE);
    }
}
