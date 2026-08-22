
import com.azure.resourcemanager.storagesync.models.ChangeDetectionMode;
import com.azure.resourcemanager.storagesync.models.TriggerChangeDetectionParameters;

/**
 * Samples for CloudEndpoints TriggerChangeDetection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/CloudEndpoints_TriggerChangeDetection.json
     */
    /**
     * Sample code: CloudEndpoints_TriggerChangeDetection.
     * 
     * @param manager Entry point to StorageSyncManager.
     */
    public static void
        cloudEndpointsTriggerChangeDetection(com.azure.resourcemanager.storagesync.StorageSyncManager manager) {
        manager.cloudEndpoints().triggerChangeDetection("SampleResourceGroup_1", "SampleStorageSyncService_1",
            "SampleSyncGroup_1", "SampleCloudEndpoint_1", new TriggerChangeDetectionParameters()
                .withDirectoryPath("NewDirectory").withChangeDetectionMode(ChangeDetectionMode.RECURSIVE),
            com.azure.core.util.Context.NONE);
    }
}
