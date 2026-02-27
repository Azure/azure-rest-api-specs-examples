
import com.azure.resourcemanager.containerregistry.models.PasswordName;
import com.azure.resourcemanager.containerregistry.models.RegenerateCredentialParameters;

/**
 * Samples for Registries RegenerateCredential.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/RegistryRegenerateCredential.json
     */
    /**
     * Sample code: RegistryRegenerateCredential.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void registryRegenerateCredential(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getRegistries().regenerateCredentialWithResponse(
            "myResourceGroup", "myRegistry", new RegenerateCredentialParameters().withName(PasswordName.PASSWORD),
            com.azure.core.util.Context.NONE);
    }
}
