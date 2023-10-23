/** Samples for JobRuns Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2023-10-01/examples/JobRuns_Get.json
     */
    /**
     * Sample code: JobRuns_Get.
     *
     * @param manager Entry point to StorageMoverManager.
     */
    public static void jobRunsGet(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager
            .jobRuns()
            .getWithResponse(
                "examples-rg",
                "examples-storageMoverName",
                "examples-projectName",
                "examples-jobDefinitionName",
                "examples-jobRunName",
                com.azure.core.util.Context.NONE);
    }
}
