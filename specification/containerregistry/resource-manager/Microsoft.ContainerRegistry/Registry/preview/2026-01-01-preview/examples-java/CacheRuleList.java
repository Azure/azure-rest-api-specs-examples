
/**
 * Samples for CacheRules List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/CacheRuleList.json
     */
    /**
     * Sample code: CacheRuleList.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void cacheRuleList(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getCacheRules().list("myResourceGroup", "myRegistry", com.azure.core.util.Context.NONE);
    }
}
