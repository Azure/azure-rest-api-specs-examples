import com.azure.core.util.Context;

/** Samples for DscNode ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listPagedDscNodesByAutomationAccountWithVersionFilter.json
     */
    /**
     * Sample code: List Paged DSC nodes by Automation Account with version filter.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listPagedDSCNodesByAutomationAccountWithVersionFilter(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .dscNodes()
            .listByAutomationAccount(
                "rg",
                "myAutomationAccount33",
                "properties/extensionHandler/any(eh: eh/version le '2.70')",
                0,
                4,
                "allpages",
                Context.NONE);
    }
}
