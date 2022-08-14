import com.azure.core.util.Context;

/** Samples for Runbook Publish. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/publishRunbook.json
     */
    /**
     * Sample code: Publish runbook draft.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void publishRunbookDraft(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.runbooks().publish("rg", "ContoseAutomationAccount", "Get-AzureVMTutorial", Context.NONE);
    }
}
