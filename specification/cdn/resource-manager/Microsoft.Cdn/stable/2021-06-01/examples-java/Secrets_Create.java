import com.azure.core.util.Context;
import com.azure.resourcemanager.cdn.fluent.models.SecretInner;
import com.azure.resourcemanager.cdn.models.CustomerCertificateParameters;
import com.azure.resourcemanager.cdn.models.ResourceReference;

/** Samples for Secrets Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Secrets_Create.json
     */
    /**
     * Sample code: Secrets_Create.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void secretsCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getSecrets()
            .create(
                "RG",
                "profile1",
                "secret1",
                new SecretInner()
                    .withParameters(
                        new CustomerCertificateParameters()
                            .withSecretSource(
                                new ResourceReference()
                                    .withId(
                                        "/subscriptions/subid/resourcegroups/RG/providers/Microsoft.KeyVault/vault/kvName/secrets/certificatename"))
                            .withSecretVersion("abcdef1234578900abcdef1234567890")
                            .withUseLatestVersion(false)),
                Context.NONE);
    }
}
