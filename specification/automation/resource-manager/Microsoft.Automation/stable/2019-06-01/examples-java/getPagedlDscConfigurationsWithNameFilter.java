import com.azure.core.util.Context;

/** Samples for DscConfiguration ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/getPagedlDscConfigurationsWithNameFilter.json
     */
    /**
     * Sample code: List Paged DSC Configurations with name filter.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listPagedDSCConfigurationsWithNameFilter(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .dscConfigurations()
            .listByAutomationAccount(
                "rg", "myAutomationAccount33", "contains(name,'server')", 0, 2, "allpages", Context.NONE);
    }
}
