
/**
 * Samples for Archives Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/ArchiveGet.json
     */
    /**
     * Sample code: ArchiveGet.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void archiveGet(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getArchives().getWithResponse("myResourceGroup", "myRegistry", "myPackageType",
            "myArchiveName", com.azure.core.util.Context.NONE);
    }
}
