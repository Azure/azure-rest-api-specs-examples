
import com.azure.resourcemanager.storagemover.models.CopyMode;
import com.azure.resourcemanager.storagemover.models.JobType;

/**
 * Samples for JobDefinitions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/JobDefinitions_CreateOrUpdate_CloudToCloud.json
     */
    /**
     * Sample code: JobDefinitions_CreateOrUpdate_CloudToCloud.
     * 
     * @param manager Entry point to StorageMoverManager.
     */
    public static void
        jobDefinitionsCreateOrUpdateCloudToCloud(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager.jobDefinitions().define("examples-jobDefinitionName")
            .withExistingProject("examples-rg", "examples-storageMoverName", "examples-projectName")
            .withCopyMode(CopyMode.ADDITIVE).withSourceName("examples-sourceEndpointName")
            .withTargetName("examples-targetEndpointName").withDescription("Example Job Definition Description")
            .withJobType(JobType.CLOUD_TO_CLOUD).withSourceSubpath("/").withTargetSubpath("/")
            .withAgentName("dummy-agent").create();
    }
}
