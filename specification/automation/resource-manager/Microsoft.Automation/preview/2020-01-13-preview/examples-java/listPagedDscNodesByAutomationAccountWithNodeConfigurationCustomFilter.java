import com.azure.core.util.Context;

/** Samples for DscNode ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listPagedDscNodesByAutomationAccountWithNodeConfigurationCustomFilter.json
     */
    /**
     * Sample code: List Paged DSC nodes by Automation Account with Node Configuration Custom filter.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listPagedDSCNodesByAutomationAccountWithNodeConfigurationCustomFilter(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .dscNodes()
            .listByAutomationAccount(
                "rg",
                "myAutomationAccount33",
                "contains(properties/nodeConfiguration/name,'SetupServer.localhost,SetupClient.localhost,$$Not$$Configured$$')",
                0,
                4,
                "allpages",
                Context.NONE);
    }
}
