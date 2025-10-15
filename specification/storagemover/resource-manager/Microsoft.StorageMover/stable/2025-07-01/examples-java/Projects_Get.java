
/**
 * Samples for Projects Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/Projects_Get.json
     */
    /**
     * Sample code: Projects_Get.
     * 
     * @param manager Entry point to StorageMoverManager.
     */
    public static void projectsGet(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager.projects().getWithResponse("examples-rg", "examples-storageMoverName", "examples-projectName",
            com.azure.core.util.Context.NONE);
    }
}
