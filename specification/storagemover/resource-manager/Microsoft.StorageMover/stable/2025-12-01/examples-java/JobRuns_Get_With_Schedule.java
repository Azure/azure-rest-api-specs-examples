
/**
 * Samples for JobRuns Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/JobRuns_Get_With_Schedule.json
     */
    /**
     * Sample code: JobRuns_Get_With_Schedule.
     * 
     * @param manager Entry point to StorageMoverManager.
     */
    public static void jobRunsGetWithSchedule(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager.jobRuns().getWithResponse("examples-rg", "examples-storageMoverName", "examples-projectName",
            "examples-jobDefinitionName", "examples-jobRunName", com.azure.core.util.Context.NONE);
    }
}
