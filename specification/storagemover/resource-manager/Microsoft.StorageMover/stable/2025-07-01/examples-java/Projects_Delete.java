
/**
 * Samples for Projects Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/Projects_Delete.json
     */
    /**
     * Sample code: Projects_Delete.
     * 
     * @param manager Entry point to StorageMoverManager.
     */
    public static void projectsDelete(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager.projects().delete("examples-rg", "examples-storageMoverName", "examples-projectName",
            com.azure.core.util.Context.NONE);
    }
}
