
import com.azure.resourcemanager.containerregistry.fluent.models.ScopeMapInner;
import java.util.Arrays;

/**
 * Samples for ScopeMaps Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ScopeMapCreate.json
     */
    /**
     * Sample code: ScopeMapCreate.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void scopeMapCreate(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getScopeMaps().create("myResourceGroup", "myRegistry", "myScopeMap",
            new ScopeMapInner().withDescription("Developer Scopes").withActions(
                Arrays.asList("repositories/myrepository/contentWrite", "repositories/myrepository/delete")),
            com.azure.core.util.Context.NONE);
    }
}
