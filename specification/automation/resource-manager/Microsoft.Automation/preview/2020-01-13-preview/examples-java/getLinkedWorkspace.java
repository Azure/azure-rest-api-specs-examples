import com.azure.core.util.Context;

/** Samples for LinkedWorkspace Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getLinkedWorkspace.json
     */
    /**
     * Sample code: Get the linked workspace of an automation account.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getTheLinkedWorkspaceOfAnAutomationAccount(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.linkedWorkspaces().getWithResponse("rg", "ContosoAutomationAccount", Context.NONE);
    }
}
