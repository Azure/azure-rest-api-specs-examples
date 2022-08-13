import com.azure.core.util.Context;

/** Samples for Runbook Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/deleteRunbook.json
     */
    /**
     * Sample code: Delete a runbook.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void deleteARunbook(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.runbooks().deleteWithResponse("rg", "ContoseAutomationAccount", "Get-AzureVMTutorial", Context.NONE);
    }
}
