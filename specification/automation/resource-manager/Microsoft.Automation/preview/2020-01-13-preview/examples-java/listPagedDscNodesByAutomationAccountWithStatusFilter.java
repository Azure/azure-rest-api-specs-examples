import com.azure.core.util.Context;

/** Samples for DscNode ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listPagedDscNodesByAutomationAccountWithStatusFilter.json
     */
    /**
     * Sample code: List Paged DSC nodes by Automation Account with node status filter.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listPagedDSCNodesByAutomationAccountWithNodeStatusFilter(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .dscNodes()
            .listByAutomationAccount(
                "rg",
                "myAutomationAccount33",
                "contains(properties/status,'Compliant,NotCompliant')",
                0,
                4,
                "allpages",
                Context.NONE);
    }
}
