/** Samples for JobDefinitions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/preview/2022-07-01-preview/examples/JobDefinitions_List.json
     */
    /**
     * Sample code: JobDefinitions_List.
     *
     * @param manager Entry point to StorageMoverManager.
     */
    public static void jobDefinitionsList(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager
            .jobDefinitions()
            .list("examples-rg", "examples-storageMoverName", "examples-projectName", com.azure.core.util.Context.NONE);
    }
}
