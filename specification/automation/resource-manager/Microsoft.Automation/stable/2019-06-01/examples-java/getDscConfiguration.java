import com.azure.core.util.Context;

/** Samples for DscConfiguration Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/getDscConfiguration.json
     */
    /**
     * Sample code: Get a DSC Configuration.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getADSCConfiguration(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.dscConfigurations().getWithResponse("rg", "myAutomationAccount33", "TemplateBasic", Context.NONE);
    }
}
