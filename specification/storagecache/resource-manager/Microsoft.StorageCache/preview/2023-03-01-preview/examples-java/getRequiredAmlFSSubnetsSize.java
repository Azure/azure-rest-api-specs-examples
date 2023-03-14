/** Samples for ResourceProvider GetRequiredAmlFSSubnetsSize. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-03-01-preview/examples/getRequiredAmlFSSubnetsSize.json
     */
    /**
     * Sample code: getRequiredAmlFilesystemSubnetsSize.
     *
     * @param manager Entry point to StorageCacheManager.
     */
    public static void getRequiredAmlFilesystemSubnetsSize(
        com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.resourceProviders().getRequiredAmlFSSubnetsSizeWithResponse(null, com.azure.core.util.Context.NONE);
    }
}
