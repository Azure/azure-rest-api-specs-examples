
import com.azure.resourcemanager.storagemover.models.JobDefinition;

/**
 * Samples for JobDefinitions Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/JobDefinitions_Update.json
     */
    /**
     * Sample code: JobDefinitions_Update.
     * 
     * @param manager Entry point to StorageMoverManager.
     */
    public static void jobDefinitionsUpdate(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        JobDefinition resource = manager.jobDefinitions().getWithResponse("examples-rg", "examples-storageMoverName",
            "examples-projectName", "examples-jobDefinitionName", com.azure.core.util.Context.NONE).getValue();
        resource.update().withDescription("Updated Job Definition Description").withAgentName("updatedAgentName")
            .apply();
    }
}
