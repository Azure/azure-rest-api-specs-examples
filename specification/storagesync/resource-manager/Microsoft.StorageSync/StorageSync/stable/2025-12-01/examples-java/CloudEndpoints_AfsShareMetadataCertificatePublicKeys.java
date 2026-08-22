
/**
 * Samples for CloudEndpoints AfsShareMetadataCertificatePublicKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/CloudEndpoints_AfsShareMetadataCertificatePublicKeys.json
     */
    /**
     * Sample code: CloudEndpoints_AfsShareMetadataCertificatePublicKeys.
     * 
     * @param manager Entry point to StorageSyncManager.
     */
    public static void cloudEndpointsAfsShareMetadataCertificatePublicKeys(
        com.azure.resourcemanager.storagesync.StorageSyncManager manager) {
        manager.cloudEndpoints().afsShareMetadataCertificatePublicKeysWithResponse("SampleResourceGroup_1",
            "SampleStorageSyncService_1", "SampleSyncGroup_1", "SampleCloudEndpoint_1",
            com.azure.core.util.Context.NONE);
    }
}
