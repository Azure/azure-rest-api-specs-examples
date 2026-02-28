
import com.azure.resourcemanager.containerregistry.models.GenerateCredentialsParameters;
import java.time.OffsetDateTime;

/**
 * Samples for Registries GenerateCredentials.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/RegistryGenerateCredentials.json
     */
    /**
     * Sample code: RegistryGenerateCredentials.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void registryGenerateCredentials(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getRegistries()
            .generateCredentials(
                "myResourceGroup", "myRegistry", new GenerateCredentialsParameters().withTokenId("fakeTokenPlaceholder")
                    .withExpiry(OffsetDateTime.parse("2020-12-31T15:59:59.0707808Z")),
                com.azure.core.util.Context.NONE);
    }
}
