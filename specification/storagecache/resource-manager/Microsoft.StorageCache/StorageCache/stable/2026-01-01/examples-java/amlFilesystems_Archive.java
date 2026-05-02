
import com.azure.resourcemanager.storagecache.models.AmlFilesystemArchiveInfo;

/**
 * Samples for AmlFilesystems Archive.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/amlFilesystems_Archive.json
     */
    /**
     * Sample code: amlFilesystems_Archive.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void amlFilesystemsArchive(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.amlFilesystems().archiveWithResponse("scgroup", "sc",
            new AmlFilesystemArchiveInfo().withFilesystemPath("/"), com.azure.core.util.Context.NONE);
    }
}
