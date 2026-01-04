
import com.azure.resourcemanager.containerregistry.models.ImportImageParameters;
import com.azure.resourcemanager.containerregistry.models.ImportMode;
import com.azure.resourcemanager.containerregistry.models.ImportSource;
import java.util.Arrays;

/**
 * Samples for Registries ImportImage.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/Registry/stable/2025-11-01/examples/
     * ImportImageFromPublicRegistry.json
     */
    /**
     * Sample code: ImportImageFromPublicRegistry.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void importImageFromPublicRegistry(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getRegistries()
            .importImage("myResourceGroup", "myRegistry", new ImportImageParameters()
                .withSource(new ImportSource().withRegistryUri("registry.hub.docker.com")
                    .withSourceImage("library/hello-world"))
                .withTargetTags(Arrays.asList("targetRepository:targetTag"))
                .withUntaggedTargetRepositories(Arrays.asList("targetRepository1")).withMode(ImportMode.FORCE),
                com.azure.core.util.Context.NONE);
    }
}
