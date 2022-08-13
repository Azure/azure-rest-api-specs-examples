import com.azure.core.util.Context;

/** Samples for DscNode ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listPagedDscNodesByAutomationAccountWithCompositeFilter.json
     */
    /**
     * Sample code: List Paged DSC nodes with filters separated by 'and'.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listPagedDSCNodesWithFiltersSeparatedByAnd(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .dscNodes()
            .listByAutomationAccount(
                "rg",
                "myAutomationAccount33",
                "properties/extensionHandler/any(eh: eh/version gt '2.70') and contains(name,'sql') and"
                    + " contains(properties/nodeConfiguration/name,'$$Not$$Configured$$')",
                0,
                10,
                "allpages",
                Context.NONE);
    }
}
