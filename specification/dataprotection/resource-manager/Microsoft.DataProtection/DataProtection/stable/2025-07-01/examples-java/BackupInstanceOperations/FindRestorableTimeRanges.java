
import com.azure.resourcemanager.dataprotection.models.AzureBackupFindRestorableTimeRangesRequest;
import com.azure.resourcemanager.dataprotection.models.RestoreSourceDataStoreType;

/**
 * Samples for RestorableTimeRanges Find.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/BackupInstanceOperations/FindRestorableTimeRanges.json
     */
    /**
     * Sample code: Find Restorable Time Ranges.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void
        findRestorableTimeRanges(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.restorableTimeRanges().findWithResponse("Blob-Backup", "ZBlobBackupVaultBVTD3", "zblobbackuptestsa58",
            new AzureBackupFindRestorableTimeRangesRequest()
                .withSourceDataStoreType(RestoreSourceDataStoreType.OPERATIONAL_STORE)
                .withStartTime("2020-10-17T23:28:17.6829685Z").withEndTime("2021-02-24T00:35:17.6829685Z"),
            com.azure.core.util.Context.NONE);
    }
}
