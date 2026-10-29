
import com.azure.resourcemanager.cdn.fluent.models.SecretInner;
import com.azure.resourcemanager.cdn.models.ResourceReference;
import com.azure.resourcemanager.cdn.models.UrlSigningKeyParameters;

/**
 * Samples for Secrets Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Secrets_CreateUrlSigningKeyType.json
     */
    /**
     * Sample code: Secrets_CreateUrlSigningKeyType.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void secretsCreateUrlSigningKeyType(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getSecrets().create("RG", "profile1", "secret1",
            new SecretInner().withParameters(new UrlSigningKeyParameters().withKeyId("fakeTokenPlaceholder")
                .withSecretSource(new ResourceReference().withId(
                    "/subscriptions/subid/resourcegroups/RG/providers/Microsoft.KeyVault/vault/kvName/secrets/urlsigningkeyname"))
                .withSecretVersion("fakeTokenPlaceholder")),
            com.azure.core.util.Context.NONE);
    }
}
