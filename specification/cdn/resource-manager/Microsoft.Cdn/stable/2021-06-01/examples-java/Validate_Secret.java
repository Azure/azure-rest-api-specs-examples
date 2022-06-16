import com.azure.core.util.Context;
import com.azure.resourcemanager.cdn.models.ResourceReference;
import com.azure.resourcemanager.cdn.models.SecretType;
import com.azure.resourcemanager.cdn.models.ValidateSecretInput;

/** Samples for Validate Secret. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Validate_Secret.json
     */
    /**
     * Sample code: Validate_Secret.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void validateSecret(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getValidates()
            .secretWithResponse(
                new ValidateSecretInput()
                    .withSecretType(SecretType.CUSTOMER_CERTIFICATE)
                    .withSecretSource(
                        new ResourceReference()
                            .withId(
                                "/subscriptions/subid/resourcegroups/RG/providers/Microsoft.KeyVault/vault/kvName/certificate/certName")),
                Context.NONE);
    }
}
