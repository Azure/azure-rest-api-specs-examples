
import com.azure.resourcemanager.storagesync.models.FeatureStatus;
import com.azure.resourcemanager.storagesync.models.LocalCacheMode;
import com.azure.resourcemanager.storagesync.models.ServerEndpoint;

/**
 * Samples for ServerEndpoints Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/ServerEndpoints_Update.json
     */
    /**
     * Sample code: ServerEndpoints_Update.
     * 
     * @param manager Entry point to StorageSyncManager.
     */
    public static void serverEndpointsUpdate(com.azure.resourcemanager.storagesync.StorageSyncManager manager) {
        ServerEndpoint resource
            = manager.serverEndpoints().getWithResponse("SampleResourceGroup_1", "SampleStorageSyncService_1",
                "SampleSyncGroup_1", "SampleServerEndpoint_1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withCloudTiering(FeatureStatus.OFF).withVolumeFreeSpacePercent(100)
            .withTierFilesOlderThanDays(0).withOfflineDataTransfer(FeatureStatus.OFF)
            .withLocalCacheMode(LocalCacheMode.UPDATE_LOCALLY_CACHED_FILES).apply();
    }
}
