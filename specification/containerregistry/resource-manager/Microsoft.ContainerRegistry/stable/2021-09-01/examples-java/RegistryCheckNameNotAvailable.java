import com.azure.core.util.Context;
import com.azure.resourcemanager.containerregistry.models.RegistryNameCheckRequest;

/** Samples for Registries CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2021-09-01/examples/RegistryCheckNameNotAvailable.json
     */
    /**
     * Sample code: RegistryCheckNameNotAvailable.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void registryCheckNameNotAvailable(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getRegistries()
            .checkNameAvailabilityWithResponse(new RegistryNameCheckRequest().withName("myRegistry"), Context.NONE);
    }
}
