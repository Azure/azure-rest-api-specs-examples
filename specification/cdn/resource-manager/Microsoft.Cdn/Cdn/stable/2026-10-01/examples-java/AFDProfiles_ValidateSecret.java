
import com.azure.resourcemanager.cdn.models.ResourceReference;
import com.azure.resourcemanager.cdn.models.SecretType;
import com.azure.resourcemanager.cdn.models.ValidateSecretInput;

/**
 * Samples for AfdProfiles ValidateSecret.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/AFDProfiles_ValidateSecret.json
     */
    /**
     * Sample code: Validate_Secret.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void validateSecret(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getAfdProfiles().validateSecretWithResponse("RG", "profile1", new ValidateSecretInput()
            .withSecretType(SecretType.CUSTOMER_CERTIFICATE)
            .withSecretSource(new ResourceReference().withId(
                "/subscriptions/subid/resourcegroups/RG/providers/Microsoft.KeyVault/vault/kvName/certificate/certName")),
            com.azure.core.util.Context.NONE);
    }
}
