
import com.azure.resourcemanager.containerregistry.fluent.models.ScopeMapInner;
import java.util.Arrays;

/**
 * Samples for ScopeMaps Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/Registry/stable/2025-11-01/examples/
     * ScopeMapCreate.json
     */
    /**
     * Sample code: ScopeMapCreate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void scopeMapCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getScopeMaps().create("myResourceGroup", "myRegistry",
            "myScopeMap",
            new ScopeMapInner().withDescription("Developer Scopes").withActions(
                Arrays.asList("repositories/myrepository/contentWrite", "repositories/myrepository/delete")),
            com.azure.core.util.Context.NONE);
    }
}
