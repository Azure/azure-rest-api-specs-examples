
import com.azure.resourcemanager.loganalytics.models.DataSourceType;

/**
 * Samples for LinkedStorageAccounts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/LinkedStorageAccountsDelete.json
     */
    /**
     * Sample code: LinkedStorageAccountsDelete.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void linkedStorageAccountsDelete(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.linkedStorageAccounts().deleteWithResponse("mms-eus", "testLinkStorageAccountsWS",
            DataSourceType.CUSTOM_LOGS, com.azure.core.util.Context.NONE);
    }
}
