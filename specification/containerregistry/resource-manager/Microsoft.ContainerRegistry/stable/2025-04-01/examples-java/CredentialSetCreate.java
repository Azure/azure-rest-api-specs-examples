
import com.azure.resourcemanager.containerregistry.fluent.models.CredentialSetInner;
import com.azure.resourcemanager.containerregistry.models.AuthCredential;
import com.azure.resourcemanager.containerregistry.models.CredentialName;
import com.azure.resourcemanager.containerregistry.models.IdentityProperties;
import com.azure.resourcemanager.containerregistry.models.ResourceIdentityType;
import java.util.Arrays;

/**
 * Samples for CredentialSets Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2025-04-01/examples/
     * CredentialSetCreate.json
     */
    /**
     * Sample code: CredentialSetCreate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void credentialSetCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getCredentialSets().create("myResourceGroup",
            "myRegistry", "myCredentialSet",
            new CredentialSetInner()
                .withIdentity(new IdentityProperties().withType(ResourceIdentityType.SYSTEM_ASSIGNED))
                .withLoginServer("docker.io")
                .withAuthCredentials(Arrays.asList(new AuthCredential().withName(CredentialName.CREDENTIAL1)
                    .withUsernameSecretIdentifier("fakeTokenPlaceholder")
                    .withPasswordSecretIdentifier("fakeTokenPlaceholder"))),
            com.azure.core.util.Context.NONE);
    }
}
