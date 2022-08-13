import com.azure.core.util.Context;

/** Samples for DeletedAutomationAccounts ListBySubscription. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2022-01-31/examples/getDeletedAutomationAccount.json
     */
    /**
     * Sample code: Get deleted automation account.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getDeletedAutomationAccount(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.deletedAutomationAccounts().listBySubscriptionWithResponse(Context.NONE);
    }
}
