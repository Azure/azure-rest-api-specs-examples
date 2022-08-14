import com.azure.core.util.Context;

/** Samples for DscConfiguration ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/getPagedDscConfigurationsWithNoFilter.json
     */
    /**
     * Sample code: List Paged DSC Configurations with no filter.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listPagedDSCConfigurationsWithNoFilter(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .dscConfigurations()
            .listByAutomationAccount("rg", "myAutomationAccount33", null, 0, 3, "allpages", Context.NONE);
    }
}
