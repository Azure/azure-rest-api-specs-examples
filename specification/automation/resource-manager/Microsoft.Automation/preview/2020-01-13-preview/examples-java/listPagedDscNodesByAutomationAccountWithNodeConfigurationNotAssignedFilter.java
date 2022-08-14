import com.azure.core.util.Context;

/** Samples for DscNode ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listPagedDscNodesByAutomationAccountWithNodeConfigurationNotAssignedFilter.json
     */
    /**
     * Sample code: List Paged DSC nodes by Automation Account where Node Configurations are not assigned filter.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listPagedDSCNodesByAutomationAccountWhereNodeConfigurationsAreNotAssignedFilter(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .dscNodes()
            .listByAutomationAccount(
                "rg",
                "myAutomationAccount33",
                "properties/nodeConfiguration/name eq ''",
                0,
                20,
                "allpages",
                Context.NONE);
    }
}
