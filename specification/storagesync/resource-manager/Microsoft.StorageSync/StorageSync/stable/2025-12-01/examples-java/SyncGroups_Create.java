
import com.azure.core.util.BinaryData;
import java.nio.charset.StandardCharsets;

/**
 * Samples for SyncGroups Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/SyncGroups_Create.json
     */
    /**
     * Sample code: SyncGroups_Create.
     * 
     * @param manager Entry point to StorageSyncManager.
     */
    public static void syncGroupsCreate(com.azure.resourcemanager.storagesync.StorageSyncManager manager) {
        manager.syncGroups().define("SampleSyncGroup_1")
            .withExistingStorageSyncService("SampleResourceGroup_1", "SampleStorageSyncService_1")
            .withProperties(BinaryData.fromBytes("{}".getBytes(StandardCharsets.UTF_8))).create();
    }
}
