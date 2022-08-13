import com.azure.core.util.Context;

/** Samples for Runbook GetContent. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/getRunbookContent.json
     */
    /**
     * Sample code: Get runbook content.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getRunbookContent(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .runbooks()
            .getContentWithResponse("rg", "ContoseAutomationAccount", "Get-AzureVMTutorial", Context.NONE);
    }
}
