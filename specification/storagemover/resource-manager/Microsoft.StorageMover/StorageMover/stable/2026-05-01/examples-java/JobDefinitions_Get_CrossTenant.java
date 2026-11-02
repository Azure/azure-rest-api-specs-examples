
/**
 * Samples for JobDefinitions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/JobDefinitions_Get_CrossTenant.json
     */
    /**
     * Sample code: JobDefinitions_Get_CrossTenant.
     * 
     * @param manager Entry point to StorageMoverManager.
     */
    public static void
        jobDefinitionsGetCrossTenant(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager.jobDefinitions().getWithResponse("examples-rg", "examples-storageMoverName", "examples-projectName",
            "examples-jobDefinitionName", com.azure.core.util.Context.NONE);
    }
}
