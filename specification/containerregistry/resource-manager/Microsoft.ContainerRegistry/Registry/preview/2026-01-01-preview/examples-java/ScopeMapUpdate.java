
import com.azure.resourcemanager.containerregistry.models.ScopeMapUpdateParameters;
import java.util.Arrays;

/**
 * Samples for ScopeMaps Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/ScopeMapUpdate.json
     */
    /**
     * Sample code: ScopeMapUpdate.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void scopeMapUpdate(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getScopeMaps().update("myResourceGroup", "myRegistry", "myScopeMap",
            new ScopeMapUpdateParameters().withDescription("Developer Scopes").withActions(
                Arrays.asList("repositories/myrepository/contentWrite", "repositories/myrepository/contentRead")),
            com.azure.core.util.Context.NONE);
    }
}
