
import com.azure.resourcemanager.cdn.fluent.models.SecretInner;
import com.azure.resourcemanager.cdn.models.CustomerCertificateParameters;
import com.azure.resourcemanager.cdn.models.ResourceReference;

/**
 * Samples for Secrets Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Secrets_Create.json
     */
    /**
     * Sample code: Secrets_Create.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void secretsCreate(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getSecrets().create("RG", "profile1", "secret1", new SecretInner()
            .withParameters(new CustomerCertificateParameters().withSecretSource(new ResourceReference().withId(
                "/subscriptions/subid/resourcegroups/RG/providers/Microsoft.KeyVault/vault/kvName/secrets/certificatename"))
                .withSecretVersion("fakeTokenPlaceholder").withUseLatestVersion(false)),
            com.azure.core.util.Context.NONE);
    }
}
