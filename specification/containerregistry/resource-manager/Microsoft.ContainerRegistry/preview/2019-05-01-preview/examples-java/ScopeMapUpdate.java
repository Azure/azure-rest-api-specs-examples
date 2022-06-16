import com.azure.core.util.Context;
import com.azure.resourcemanager.containerregistry.models.ScopeMapUpdateParameters;
import java.util.Arrays;

/** Samples for ScopeMaps Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-05-01-preview/examples/ScopeMapUpdate.json
     */
    /**
     * Sample code: ScopeMapUpdate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void scopeMapUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getScopeMaps()
            .update(
                "myResourceGroup",
                "myRegistry",
                "myScopeMap",
                new ScopeMapUpdateParameters()
                    .withDescription("Developer Scopes")
                    .withActions(
                        Arrays
                            .asList("repositories/myrepository/contentWrite", "repositories/myrepository/contentRead")),
                Context.NONE);
    }
}
