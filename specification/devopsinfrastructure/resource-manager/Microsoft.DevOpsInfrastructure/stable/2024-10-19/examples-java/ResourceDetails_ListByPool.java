
/**
 * Samples for ResourceDetails ListByPool.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-10-19/ResourceDetails_ListByPool.json
     */
    /**
     * Sample code: ResourceDetails_ListByPool.
     * 
     * @param manager Entry point to DevOpsInfrastructureManager.
     */
    public static void
        resourceDetailsListByPool(com.azure.resourcemanager.devopsinfrastructure.DevOpsInfrastructureManager manager) {
        manager.resourceDetails().listByPool("my-resource-group", "my-dev-ops-pool", com.azure.core.util.Context.NONE);
    }
}
