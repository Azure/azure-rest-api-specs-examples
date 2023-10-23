/** Samples for Projects Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2023-10-01/examples/Projects_Get.json
     */
    /**
     * Sample code: Projects_Get.
     *
     * @param manager Entry point to StorageMoverManager.
     */
    public static void projectsGet(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager
            .projects()
            .getWithResponse(
                "examples-rg", "examples-storageMoverName", "examples-projectName", com.azure.core.util.Context.NONE);
    }
}
