import com.azure.resourcemanager.loganalytics.models.StorageAccount;
import java.util.Arrays;

/** Samples for StorageInsightConfigs CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/StorageInsightsCreateOrUpdate.json
     */
    /**
     * Sample code: StorageInsightsCreate.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void storageInsightsCreate(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager
            .storageInsightConfigs()
            .define("AzTestSI1110")
            .withExistingWorkspace("OIAutoRest5123", "aztest5048")
            .withContainers(Arrays.asList("wad-iis-logfiles"))
            .withTables(Arrays.asList("WADWindowsEventLogsTable", "LinuxSyslogVer2v0"))
            .withStorageAccount(
                new StorageAccount()
                    .withId(
                        "/subscriptions/00000000-0000-0000-0000-000000000005/resourcegroups/OIAutoRest6987/providers/microsoft.storage/storageaccounts/AzTestFakeSA9945")
                    .withKey("1234"))
            .create();
    }
}
