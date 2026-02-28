
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
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void registryCheckNameAvailable(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getRegistries().checkNameAvailabilityWithResponse(
            new RegistryNameCheckRequest().withName("myRegistry"), com.azure.core.util.Context.NONE);
    }
}
