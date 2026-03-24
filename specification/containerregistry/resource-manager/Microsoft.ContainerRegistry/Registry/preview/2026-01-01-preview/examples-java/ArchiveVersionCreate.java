
/**
 * Samples for ArchiveVersions Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/ArchiveVersionCreate.json
     */
    /**
     * Sample code: ArchiveVersionCreate.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void
        archiveVersionCreate(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getArchiveVersions().create("myResourceGroup", "myRegistry", "rpm", "myArchiveName",
            "myArchiveVersionName", com.azure.core.util.Context.NONE);
    }
}
