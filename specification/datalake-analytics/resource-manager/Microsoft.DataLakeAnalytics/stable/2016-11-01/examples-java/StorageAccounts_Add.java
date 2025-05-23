
import com.azure.resourcemanager.datalakeanalytics.models.AddStorageAccountParameters;

/**
 * Samples for StorageAccounts Add.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/
     * StorageAccounts_Add.json
     */
    /**
     * Sample code: Adds an Azure Storage account.
     * 
     * @param manager Entry point to DataLakeAnalyticsManager.
     */
    public static void
        addsAnAzureStorageAccount(com.azure.resourcemanager.datalakeanalytics.DataLakeAnalyticsManager manager) {
        manager.storageAccounts().addWithResponse("contosorg", "contosoadla", "test_storage",
            new AddStorageAccountParameters().withAccessKey("fakeTokenPlaceholder").withSuffix("test_suffix"),
            com.azure.core.util.Context.NONE);
    }
}
