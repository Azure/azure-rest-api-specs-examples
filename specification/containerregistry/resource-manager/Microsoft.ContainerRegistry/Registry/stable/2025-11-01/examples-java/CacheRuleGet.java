
/**
 * Samples for CacheRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/CacheRuleGet.json
     */
    /**
     * Sample code: CacheRuleGet.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void cacheRuleGet(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getCacheRules().getWithResponse("myResourceGroup", "myRegistry", "myCacheRule",
            com.azure.core.util.Context.NONE);
    }
}
