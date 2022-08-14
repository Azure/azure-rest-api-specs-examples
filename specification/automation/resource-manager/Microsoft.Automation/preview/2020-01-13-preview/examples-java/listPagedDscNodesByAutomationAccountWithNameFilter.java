import com.azure.core.util.Context;

/** Samples for DscNode ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listPagedDscNodesByAutomationAccountWithNameFilter.json
     */
    /**
     * Sample code: List Paged DSC nodes by Automation Account with name filter.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listPagedDSCNodesByAutomationAccountWithNameFilter(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .dscNodes()
            .listByAutomationAccount(
                "rg", "myAutomationAccount33", "contains('DSCCOMP',name)", 0, 6, "allpages", Context.NONE);
    }
}
