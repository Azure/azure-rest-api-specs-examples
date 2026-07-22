
import com.azure.resourcemanager.containerregistry.models.ImportImageParameters;
import com.azure.resourcemanager.containerregistry.models.ImportMode;
import com.azure.resourcemanager.containerregistry.models.ImportSource;
import java.util.Arrays;

/**
 * Samples for Registries ImportImage.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ImportImageFromPublicRegistry.json
     */
    /**
     * Sample code: ImportImageFromPublicRegistry.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void
        importImageFromPublicRegistry(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getRegistries().importImage("myResourceGroup", "myRegistry", new ImportImageParameters()
            .withSource(
                new ImportSource().withRegistryUri("registry.hub.docker.com").withSourceImage("library/hello-world"))
            .withTargetTags(Arrays.asList("targetRepository:targetTag"))
            .withUntaggedTargetRepositories(Arrays.asList("targetRepository1")).withMode(ImportMode.FORCE),
            com.azure.core.util.Context.NONE);
    }
}
