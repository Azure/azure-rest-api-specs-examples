import com.azure.core.util.Context;

/** Samples for BlobContainers ObjectLevelWorm. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/ObjectLevelWormContainerMigration.json
     */
    /**
     * Sample code: VersionLevelWormContainerMigration.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void versionLevelWormContainerMigration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getBlobContainers()
            .objectLevelWorm("res1782", "sto7069", "container6397", Context.NONE);
    }
}
