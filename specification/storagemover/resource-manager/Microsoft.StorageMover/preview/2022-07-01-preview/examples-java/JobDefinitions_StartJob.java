/** Samples for JobDefinitions StartJob. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/preview/2022-07-01-preview/examples/JobDefinitions_StartJob.json
     */
    /**
     * Sample code: JobDefinitions_StartJob.
     *
     * @param manager Entry point to StorageMoverManager.
     */
    public static void jobDefinitionsStartJob(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager
            .jobDefinitions()
            .startJobWithResponse(
                "examples-rg",
                "examples-storageMoverName",
                "examples-projectName",
                "examples-jobDefinitionName",
                com.azure.core.util.Context.NONE);
    }
}
