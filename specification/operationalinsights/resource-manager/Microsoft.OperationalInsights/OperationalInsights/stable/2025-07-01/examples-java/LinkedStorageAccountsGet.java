
import com.azure.resourcemanager.loganalytics.models.DataSourceType;

/**
 * Samples for LinkedStorageAccounts Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/LinkedStorageAccountsGet.json
     */
    /**
     * Sample code: LinkedStorageAccountsGet.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void linkedStorageAccountsGet(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.linkedStorageAccounts().getWithResponse("mms-eus", "testLinkStorageAccountsWS",
            DataSourceType.CUSTOM_LOGS, com.azure.core.util.Context.NONE);
    }
}
