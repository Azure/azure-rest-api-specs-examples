
/**
 * Samples for JobDefinitions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/JobDefinitions_Get_CrossTenant_Mirror.json
     */
    /**
     * Sample code: JobDefinitions_Get_CrossTenant_Mirror.
     * 
     * @param manager Entry point to StorageMoverManager.
     */
    public static void
        jobDefinitionsGetCrossTenantMirror(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager.jobDefinitions().getWithResponse("partner-rg", "partner-storageMover", "partner-projectName",
            "examples-jobDefinitionName", com.azure.core.util.Context.NONE);
    }
}
