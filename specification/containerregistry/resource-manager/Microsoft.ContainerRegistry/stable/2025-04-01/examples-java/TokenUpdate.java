
import com.azure.resourcemanager.containerregistry.models.TokenCertificate;
import com.azure.resourcemanager.containerregistry.models.TokenCertificateName;
import com.azure.resourcemanager.containerregistry.models.TokenCredentialsProperties;
import com.azure.resourcemanager.containerregistry.models.TokenUpdateParameters;
import java.util.Arrays;

/**
 * Samples for Tokens Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2025-04-01/examples/
     * TokenUpdate.json
     */
    /**
     * Sample code: TokenUpdate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void tokenUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getTokens()
            .update("myResourceGroup", "myRegistry", "myToken", new TokenUpdateParameters().withScopeMapId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/scopeMaps/myNewScopeMap")
                .withCredentials(new TokenCredentialsProperties().withCertificates(Arrays.asList(new TokenCertificate()
                    .withName(TokenCertificateName.CERTIFICATE1).withEncodedPemCertificate("fakeTokenPlaceholder")))),
                com.azure.core.util.Context.NONE);
    }
}
