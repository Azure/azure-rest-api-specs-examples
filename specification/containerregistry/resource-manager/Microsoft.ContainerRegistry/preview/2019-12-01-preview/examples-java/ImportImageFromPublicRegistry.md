Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.containerregistry.models.ImportImageParameters;
import com.azure.resourcemanager.containerregistry.models.ImportMode;
import com.azure.resourcemanager.containerregistry.models.ImportSource;
import java.util.Arrays;

/** Samples for Registries ImportImage. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-12-01-preview/examples/ImportImageFromPublicRegistry.json
     */
    /**
     * Sample code: ImportImageFromPublicRegistry.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void importImageFromPublicRegistry(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getRegistries()
            .importImage(
                "myResourceGroup",
                "myRegistry",
                new ImportImageParameters()
                    .withSource(
                        new ImportSource()
                            .withRegistryUri("registry.hub.docker.com")
                            .withSourceImage("library/hello-world"))
                    .withTargetTags(Arrays.asList("targetRepository:targetTag"))
                    .withUntaggedTargetRepositories(Arrays.asList("targetRepository1"))
                    .withMode(ImportMode.FORCE),
                Context.NONE);
    }
}
```
