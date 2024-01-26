
import com.azure.resourcemanager.containerregistry.models.ScopeMapUpdateParameters;
import java.util.Arrays;

/** Samples for ScopeMaps Update. */
public final class Main {
    /*
     * x-ms-original-file:
     * mgmt_containerregistry_add_readonly/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/
     * stable/2023-07-01/examples/ScopeMapUpdate.json
     */
    /**
     * Sample code: ScopeMapUpdate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void scopeMapUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getScopeMaps().update("myResourceGroup", "myRegistry",
            "myScopeMap",
            new ScopeMapUpdateParameters().withDescription("Developer Scopes").withActions(
                Arrays.asList("repositories/myrepository/contentWrite", "repositories/myrepository/contentRead")),
            com.azure.core.util.Context.NONE);
    }
}
