import com.azure.core.util.Context;

/** Samples for RunbookDraft Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/getRunbookDraft.json
     */
    /**
     * Sample code: Get runbook draft.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getRunbookDraft(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.runbookDrafts().getWithResponse("rg", "ContoseAutomationAccount", "Get-AzureVMTutorial", Context.NONE);
    }
}
