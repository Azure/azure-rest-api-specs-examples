
/**
 * Samples for JobDefinitions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2024-07-01/examples/JobDefinitions_Get.
     * json
     */
    /**
     * Sample code: JobDefinitions_Get.
     * 
     * @param manager Entry point to StorageMoverManager.
     */
    public static void jobDefinitionsGet(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager.jobDefinitions().getWithResponse("examples-rg", "examples-storageMoverName", "examples-projectName",
            "examples-jobDefinitionName", com.azure.core.util.Context.NONE);
    }
}
