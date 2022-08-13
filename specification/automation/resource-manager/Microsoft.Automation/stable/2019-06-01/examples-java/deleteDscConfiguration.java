import com.azure.core.util.Context;

/** Samples for DscConfiguration Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/deleteDscConfiguration.json
     */
    /**
     * Sample code: Delete DSC Configuration.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void deleteDSCConfiguration(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.dscConfigurations().deleteWithResponse("rg", "myAutomationAccount33", "TemplateBasic", Context.NONE);
    }
}
