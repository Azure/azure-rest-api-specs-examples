import com.azure.core.util.Context;

/** Samples for Runbook Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/getRunbook.json
     */
    /**
     * Sample code: Get runbook.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getRunbook(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.runbooks().getWithResponse("rg", "ContoseAutomationAccount", "Get-AzureVMTutorial", Context.NONE);
    }
}
