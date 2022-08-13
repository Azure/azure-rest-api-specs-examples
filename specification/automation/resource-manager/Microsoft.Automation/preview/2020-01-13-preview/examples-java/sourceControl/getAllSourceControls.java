import com.azure.core.util.Context;

/** Samples for SourceControl ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/sourceControl/getAllSourceControls.json
     */
    /**
     * Sample code: List sourceControls.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listSourceControls(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.sourceControls().listByAutomationAccount("rg", "sampleAccount9", null, Context.NONE);
    }
}
