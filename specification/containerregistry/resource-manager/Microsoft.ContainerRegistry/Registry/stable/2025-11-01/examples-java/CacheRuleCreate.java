
import com.azure.resourcemanager.containerregistry.fluent.models.CacheRuleInner;

/**
 * Samples for CacheRules Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/CacheRuleCreate.json
     */
    /**
     * Sample code: CacheRuleCreate.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void cacheRuleCreate(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getCacheRules().create("myResourceGroup", "myRegistry", "myCacheRule",
            new CacheRuleInner().withCredentialSetResourceId("fakeTokenPlaceholder")
                .withSourceRepository("docker.io/library/hello-world")
                .withTargetRepository("cached-docker-hub/hello-world"),
            com.azure.core.util.Context.NONE);
    }
}
