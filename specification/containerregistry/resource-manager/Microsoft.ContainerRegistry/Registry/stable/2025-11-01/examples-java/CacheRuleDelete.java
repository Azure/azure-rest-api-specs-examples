
/**
 * Samples for CacheRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/CacheRuleDelete.json
     */
    /**
     * Sample code: CacheRuleDelete.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void cacheRuleDelete(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getCacheRules().delete("myResourceGroup", "myRegistry", "myCacheRule",
            com.azure.core.util.Context.NONE);
    }
}
