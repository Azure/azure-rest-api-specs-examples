
import com.azure.resourcemanager.containerregistry.models.PasswordName;
import com.azure.resourcemanager.containerregistry.models.RegenerateCredentialParameters;

/**
 * Samples for Registries RegenerateCredential.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/RegistryRegenerateCredential.json
     */
    /**
     * Sample code: RegistryRegenerateCredential.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void
        registryRegenerateCredential(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getRegistries().regenerateCredentialWithResponse("myResourceGroup", "myRegistry",
            new RegenerateCredentialParameters().withName(PasswordName.PASSWORD), com.azure.core.util.Context.NONE);
    }
}
