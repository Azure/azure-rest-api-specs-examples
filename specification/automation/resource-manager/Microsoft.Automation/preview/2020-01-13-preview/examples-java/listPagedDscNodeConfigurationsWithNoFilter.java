import com.azure.core.util.Context;

/** Samples for DscNodeConfiguration ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listPagedDscNodeConfigurationsWithNoFilter.json
     */
    /**
     * Sample code: List Paged DSC node configurations by Automation Account with no filter.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listPagedDSCNodeConfigurationsByAutomationAccountWithNoFilter(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .dscNodeConfigurations()
            .listByAutomationAccount("rg", "myAutomationAccount33", null, 0, 4, "allpages", Context.NONE);
    }
}
