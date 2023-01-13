import com.azure.resourcemanager.datalakeanalytics.models.UpdateStorageAccountParameters;

/** Samples for StorageAccounts Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/StorageAccounts_Update.json
     */
    /**
     * Sample code: Replaces Azure Storage blob account details.
     *
     * @param manager Entry point to DataLakeAnalyticsManager.
     */
    public static void replacesAzureStorageBlobAccountDetails(
        com.azure.resourcemanager.datalakeanalytics.DataLakeAnalyticsManager manager) {
        manager
            .storageAccounts()
            .updateWithResponse(
                "contosorg",
                "contosoadla",
                "test_storage",
                new UpdateStorageAccountParameters().withAccessKey("fakeTokenPlaceholder").withSuffix("test_suffix"),
                com.azure.core.util.Context.NONE);
    }
}
