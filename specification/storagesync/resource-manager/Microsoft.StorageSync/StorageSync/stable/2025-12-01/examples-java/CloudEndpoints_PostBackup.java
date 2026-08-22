
import com.azure.resourcemanager.storagesync.models.BackupRequest;

/**
 * Samples for CloudEndpoints PostBackup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/CloudEndpoints_PostBackup.json
     */
    /**
     * Sample code: CloudEndpoints_PostBackup.
     * 
     * @param manager Entry point to StorageSyncManager.
     */
    public static void cloudEndpointsPostBackup(com.azure.resourcemanager.storagesync.StorageSyncManager manager) {
        manager.cloudEndpoints().postBackup("SampleResourceGroup_1", "SampleStorageSyncService_1", "SampleSyncGroup_1",
            "SampleCloudEndpoint_1",
            new BackupRequest()
                .withAzureFileShare("https://sampleserver.file.core.test-cint.azure-test.net/sampleFileShare"),
            com.azure.core.util.Context.NONE);
    }
}
