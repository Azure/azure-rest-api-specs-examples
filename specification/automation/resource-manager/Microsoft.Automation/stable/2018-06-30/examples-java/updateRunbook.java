
import com.azure.resourcemanager.automation.models.Runbook;

/**
 * Samples for Runbook Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/updateRunbook.json
     */
    /**
     * Sample code: Update runbook.
     * 
     * @param manager Entry point to AutomationManager.
     */
    public static void updateRunbook(com.azure.resourcemanager.automation.AutomationManager manager) {
        Runbook resource = manager.runbooks()
            .getWithResponse("rg", "ContoseAutomationAccount", "Get-AzureVMTutorial", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withDescription("Updated Description of the Runbook").withLogVerbose(false)
            .withLogProgress(true).withLogActivityTrace(1).apply();
    }
}
