
/**
 * Samples for RunbookDraft UndoEdit.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/
     * undoDraftEditToLastKnownPublishedState.json
     */
    /**
     * Sample code: Undo draft edit to last known published state.
     * 
     * @param manager Entry point to AutomationManager.
     */
    public static void
        undoDraftEditToLastKnownPublishedState(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.runbookDrafts().undoEditWithResponse("rg", "ContoseAutomationAccount", "Get-AzureVMTutorial",
            com.azure.core.util.Context.NONE);
    }
}
