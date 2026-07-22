
import com.azure.resourcemanager.containerregistry.models.ImportImageParameters;
import com.azure.resourcemanager.containerregistry.models.ImportMode;
import com.azure.resourcemanager.containerregistry.models.ImportSource;
import java.util.Arrays;

/**
 * Samples for Registries ImportImage.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ImportImageByManifestDigest.json
     */
    /**
     * Sample code: ImportImageByManifestDigest.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void
        importImageByManifestDigest(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getRegistries().importImage("myResourceGroup", "myRegistry",
            new ImportImageParameters().withSource(new ImportSource().withResourceId(
                "/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/sourceResourceGroup/providers/Microsoft.ContainerRegistry/registries/sourceRegistry")
                .withSourceImage(
                    "sourceRepository@sha256:0000000000000000000000000000000000000000000000000000000000000000"))
                .withTargetTags(Arrays.asList("targetRepository:targetTag"))
                .withUntaggedTargetRepositories(Arrays.asList("targetRepository1")).withMode(ImportMode.FORCE),
            com.azure.core.util.Context.NONE);
    }
}
