
/**
 * Samples for JobDefinitions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/JobDefinitions_Get_With_Schedule.json
     */
    /**
     * Sample code: JobDefinitions_Get_With_Schedule.
     * 
     * @param manager Entry point to StorageMoverManager.
     */
    public static void
        jobDefinitionsGetWithSchedule(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager.jobDefinitions().getWithResponse("examples-rg", "examples-storageMoverName", "examples-projectName",
            "examples-jobDefinitionName", com.azure.core.util.Context.NONE);
    }
}
