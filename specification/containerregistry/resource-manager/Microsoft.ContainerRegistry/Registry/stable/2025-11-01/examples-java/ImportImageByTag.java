
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
     * ImportImageByTag.json
     */
    /**
     * Sample code: ImportImageByTag.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void importImageByTag(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getRegistries().importImage("myResourceGroup",
            "myRegistry",
            new ImportImageParameters().withSource(new ImportSource().withResourceId(
                "/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/sourceResourceGroup/providers/Microsoft.ContainerRegistry/registries/sourceRegistry")
                .withSourceImage("sourceRepository:sourceTag"))
                .withTargetTags(Arrays.asList("targetRepository:targetTag"))
                .withUntaggedTargetRepositories(Arrays.asList("targetRepository1")).withMode(ImportMode.FORCE),
            com.azure.core.util.Context.NONE);
    }
}
