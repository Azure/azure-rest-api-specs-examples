
/**
 * Samples for ArchiveVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/ArchiveVersionDelete.json
     */
    /**
     * Sample code: ArchiveVersionDelete.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void
        archiveVersionDelete(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getArchiveVersions().delete("myResourceGroup", "myRegistry", "myPackageType",
            "myArchiveName", "myArchiveVersionName", com.azure.core.util.Context.NONE);
    }
}
