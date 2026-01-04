
import com.azure.resourcemanager.containerregistry.fluent.models.CacheRuleInner;

/**
 * Samples for CacheRules Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/Registry/stable/2025-11-01/examples/
     * CacheRuleCreate.json
     */
    /**
     * Sample code: CacheRuleCreate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cacheRuleCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getCacheRules().create("myResourceGroup", "myRegistry",
            "myCacheRule",
            new CacheRuleInner().withCredentialSetResourceId("fakeTokenPlaceholder")
                .withSourceRepository("docker.io/library/hello-world")
                .withTargetRepository("cached-docker-hub/hello-world"),
            com.azure.core.util.Context.NONE);
    }
}
