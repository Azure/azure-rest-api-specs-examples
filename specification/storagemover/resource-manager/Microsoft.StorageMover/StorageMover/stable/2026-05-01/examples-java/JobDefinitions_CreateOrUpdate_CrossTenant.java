
import com.azure.resourcemanager.storagemover.models.CopyMode;
import com.azure.resourcemanager.storagemover.models.JobType;

/**
 * Samples for JobDefinitions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/JobDefinitions_CreateOrUpdate_CrossTenant.json
     */
    /**
     * Sample code: JobDefinitions_CreateOrUpdate_CrossTenant.
     * 
     * @param manager Entry point to StorageMoverManager.
     */
    public static void
        jobDefinitionsCreateOrUpdateCrossTenant(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager.jobDefinitions().define("examples-jobDefinitionName")
            .withExistingProject("examples-rg", "examples-storageMoverName", "examples-projectName")
            .withCopyMode(CopyMode.ADDITIVE).withSourceName("examples-sourceEndpointName")
            .withTargetName("partner-targetEndpoint")
            .withDescription(
                "Cross-tenant Blob-to-Blob copy job. JobDefinition lives in the source (host) tenant alongside the local source endpoint; the target endpoint lives in a partner (guest) tenant.")
            .withJobType(JobType.CLOUD_TO_CLOUD).withSourceSubpath("/").withTargetSubpath("/")
            .withAgentName("migration-agent").withIsCrossTenantJob(true)
            .withCrossTenantEndpointTenantId("11111111-2222-3333-4444-555555555555")
            .withCrossTenantEndpointResourceId(
                "/subscriptions/0a2b3c4d-5e6f-7081-92a3-b4c5d6e7f809/resourceGroups/partner-rg/providers/Microsoft.StorageMover/storageMovers/partner-storageMover/endpoints/partner-targetEndpoint")
            .create();
    }
}
