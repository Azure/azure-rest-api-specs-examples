
/**
 * Samples for BlobContainers ObjectLevelWorm.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/ObjectLevelWormContainerMigration.json
     */
    /**
     * Sample code: VersionLevelWormContainerMigration.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void versionLevelWormContainerMigration(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobContainers().objectLevelWorm("res1782", "sto7069", "container6397",
            com.azure.core.util.Context.NONE);
    }
}
