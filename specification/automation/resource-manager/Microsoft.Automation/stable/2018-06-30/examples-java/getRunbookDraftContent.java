import com.azure.core.util.Context;

/** Samples for RunbookDraft GetContent. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/getRunbookDraftContent.json
     */
    /**
     * Sample code: Get runbook draft content.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getRunbookDraftContent(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .runbookDrafts()
            .getContentWithResponse("rg", "ContoseAutomationAccount", "Get-AzureVMTutorial", Context.NONE);
    }
}
