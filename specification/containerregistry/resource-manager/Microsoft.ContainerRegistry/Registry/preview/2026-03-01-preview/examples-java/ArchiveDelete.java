
/**
 * Samples for Archives Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ArchiveDelete.json
     */
    /**
     * Sample code: ArchiveDelete.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void archiveDelete(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getArchives().delete("myResourceGroup", "myRegistry", "myPackageType", "myArchiveName",
            com.azure.core.util.Context.NONE);
    }
}
