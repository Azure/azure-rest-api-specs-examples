import com.azure.core.util.Context;

/** Samples for DscConfiguration ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/getAllDscConfigurations.json
     */
    /**
     * Sample code: Get DSC Configuration.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getDSCConfiguration(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .dscConfigurations()
            .listByAutomationAccount("rg", "myAutomationAccount33", null, null, null, null, Context.NONE);
    }
}
