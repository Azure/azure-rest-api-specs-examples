import com.azure.core.util.Context;

/** Samples for DscNode ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listPagedDscNodesByAutomationAccountWithNoFilter.json
     */
    /**
     * Sample code: List Paged DSC nodes by Automation Account with no filters.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listPagedDSCNodesByAutomationAccountWithNoFilters(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.dscNodes().listByAutomationAccount("rg", "myAutomationAccount33", null, 0, 2, "allpages", Context.NONE);
    }
}
