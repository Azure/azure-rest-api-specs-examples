
import com.azure.resourcemanager.containerregistry.models.RegistryNameCheckRequest;

/**
 * Samples for Registries CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/RegistryCheckNameAvailable.json
     */
    /**
     * Sample code: RegistryCheckNameAvailable.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void
        registryCheckNameAvailable(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getRegistries().checkNameAvailabilityWithResponse(
            new RegistryNameCheckRequest().withName("myRegistry"), com.azure.core.util.Context.NONE);
    }
}
