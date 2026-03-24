
import com.azure.resourcemanager.containerregistry.models.ArchiveUpdateParameters;

/**
 * Samples for Archives Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/ArchiveUpdate.json
     */
    /**
     * Sample code: ArchiveUpdate.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void archiveUpdate(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getArchives().updateWithResponse("myResourceGroup", "myRegistry", "myPackageType",
            "myArchiveName", new ArchiveUpdateParameters().withPublishedVersion("string"),
            com.azure.core.util.Context.NONE);
    }
}
