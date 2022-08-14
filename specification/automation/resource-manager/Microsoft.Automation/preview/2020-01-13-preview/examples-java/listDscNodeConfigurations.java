import com.azure.core.util.Context;

/** Samples for DscNodeConfiguration ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listDscNodeConfigurations.json
     */
    /**
     * Sample code: List DSC node configurations by Automation Account.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listDSCNodeConfigurationsByAutomationAccount(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .dscNodeConfigurations()
            .listByAutomationAccount("rg", "myAutomationAccount33", null, null, null, null, Context.NONE);
    }
}
