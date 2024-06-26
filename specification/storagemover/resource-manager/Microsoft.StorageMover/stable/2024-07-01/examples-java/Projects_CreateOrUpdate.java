
/**
 * Samples for Projects CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2024-07-01/examples/
     * Projects_CreateOrUpdate.json
     */
    /**
     * Sample code: Projects_CreateOrUpdate.
     * 
     * @param manager Entry point to StorageMoverManager.
     */
    public static void projectsCreateOrUpdate(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager.projects().define("examples-projectName")
            .withExistingStorageMover("examples-rg", "examples-storageMoverName")
            .withDescription("Example Project Description").create();
    }
}
