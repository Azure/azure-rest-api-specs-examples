
/**
 * Samples for Archives List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/ArchiveList.json
     */
    /**
     * Sample code: ArchiveList.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void archiveList(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getArchives().list("myResourceGroup", "myRegistry", "myPackageType",
            com.azure.core.util.Context.NONE);
    }
}
