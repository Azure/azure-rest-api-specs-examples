import com.azure.resourcemanager.loganalytics.models.DataSourceType;

/** Samples for LinkedStorageAccounts Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/LinkedStorageAccountsDelete.json
     */
    /**
     * Sample code: LinkedStorageAccountsDelete.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void linkedStorageAccountsDelete(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager
            .linkedStorageAccounts()
            .deleteWithResponse(
                "mms-eus", "testLinkStorageAccountsWS", DataSourceType.CUSTOM_LOGS, com.azure.core.util.Context.NONE);
    }
}
