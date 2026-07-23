
/**
 * Samples for ArchiveVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ArchiveVersionGet.json
     */
    /**
     * Sample code: ArchiveVersionGet.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void archiveVersionGet(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getArchiveVersions().getWithResponse("myResourceGroup", "myRegistry", "rpm",
            "myArchiveName", "myArchiveVersionName", com.azure.core.util.Context.NONE);
    }
}
