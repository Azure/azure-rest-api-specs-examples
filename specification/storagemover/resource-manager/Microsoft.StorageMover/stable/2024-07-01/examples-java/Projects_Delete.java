
/**
 * Samples for Projects Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2024-07-01/examples/Projects_Delete.
     * json
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
