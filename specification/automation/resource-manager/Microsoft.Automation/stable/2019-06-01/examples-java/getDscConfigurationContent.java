import com.azure.core.util.Context;

/** Samples for DscConfiguration GetContent. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/getDscConfigurationContent.json
     */
    /**
     * Sample code: Get DSC Configuration content.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getDSCConfigurationContent(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.dscConfigurations().getContentWithResponse("rg", "myAutomationAccount33", "ConfigName", Context.NONE);
    }
}
