
import com.azure.resourcemanager.appplatform.fluent.models.ContainerRegistryResourceInner;
import com.azure.resourcemanager.appplatform.models.ContainerRegistryBasicCredentials;
import com.azure.resourcemanager.appplatform.models.ContainerRegistryProperties;

/**
 * Samples for ContainerRegistries CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * ContainerRegistries_CreateOrUpdate.json
     */
    /**
     * Sample code: ContainerRegistries_CreateOrUpdate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void containerRegistriesCreateOrUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getContainerRegistries().createOrUpdate("myResourceGroup",
            "my-service", "my-container-registry",
            new ContainerRegistryResourceInner().withProperties(
                new ContainerRegistryProperties().withCredentials(new ContainerRegistryBasicCredentials()
                    .withServer("myServer").withUsername("myUsername").withPassword("fakeTokenPlaceholder"))),
            com.azure.core.util.Context.NONE);
    }
}
