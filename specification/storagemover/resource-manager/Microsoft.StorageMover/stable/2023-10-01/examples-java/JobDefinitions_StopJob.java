/** Samples for JobDefinitions StopJob. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2023-10-01/examples/JobDefinitions_StopJob.json
     */
    /**
     * Sample code: JobDefinitions_StopJob.
     *
     * @param manager Entry point to StorageMoverManager.
     */
    public static void jobDefinitionsStopJob(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager
            .jobDefinitions()
            .stopJobWithResponse(
                "examples-rg",
                "examples-storageMoverName",
                "examples-projectName",
                "examples-jobDefinitionName",
                com.azure.core.util.Context.NONE);
    }
}
