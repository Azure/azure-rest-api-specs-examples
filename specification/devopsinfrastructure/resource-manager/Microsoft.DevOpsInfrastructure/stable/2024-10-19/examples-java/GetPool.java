
/**
 * Samples for Pools GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-10-19/GetPool.json
     */
    /**
     * Sample code: Pools_Get.
     * 
     * @param manager Entry point to DevOpsInfrastructureManager.
     */
    public static void poolsGet(com.azure.resourcemanager.devopsinfrastructure.DevOpsInfrastructureManager manager) {
        manager.pools().getByResourceGroupWithResponse("rg", "pool", com.azure.core.util.Context.NONE);
    }
}
