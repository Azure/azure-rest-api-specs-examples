
import com.azure.resourcemanager.storagesync.models.RecallActionParameters;

/**
 * Samples for ServerEndpoints RecallAction.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/ServerEndpoints_Recall.json
     */
    /**
     * Sample code: ServerEndpoints_recallAction.
     * 
     * @param manager Entry point to StorageSyncManager.
     */
    public static void serverEndpointsRecallAction(com.azure.resourcemanager.storagesync.StorageSyncManager manager) {
        manager.serverEndpoints().recallAction("SampleResourceGroup_1", "SampleStorageSyncService_1",
            "SampleSyncGroup_1", "SampleServerEndpoint_1",
            new RecallActionParameters().withPattern("").withRecallPath(""), com.azure.core.util.Context.NONE);
    }
}
