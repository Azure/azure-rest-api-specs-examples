/** Samples for JobRuns List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2023-10-01/examples/JobRuns_List.json
     */
    /**
     * Sample code: JobRuns_List.
     *
     * @param manager Entry point to StorageMoverManager.
     */
    public static void jobRunsList(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager
            .jobRuns()
            .list(
                "examples-rg",
                "examples-storageMoverName",
                "examples-projectName",
                "examples-jobDefinitionName",
                com.azure.core.util.Context.NONE);
    }
}
