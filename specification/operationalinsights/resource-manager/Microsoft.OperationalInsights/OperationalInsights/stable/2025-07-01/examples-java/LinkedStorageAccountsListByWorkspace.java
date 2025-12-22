
/**
 * Samples for LinkedStorageAccounts ListByWorkspace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/LinkedStorageAccountsListByWorkspace.json
     */
    /**
     * Sample code: Gets list of linked storage accounts on a workspace.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void getsListOfLinkedStorageAccountsOnAWorkspace(
        com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.linkedStorageAccounts().listByWorkspace("mms-eus", "testLinkStorageAccountsWS",
            com.azure.core.util.Context.NONE);
    }
}
