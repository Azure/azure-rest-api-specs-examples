
import com.azure.resourcemanager.cognitiveservices.models.AccountProperties;
import com.azure.resourcemanager.cognitiveservices.models.CertificateBundleReference;
import com.azure.resourcemanager.cognitiveservices.models.Identity;
import com.azure.resourcemanager.cognitiveservices.models.ResourceIdentityType;
import com.azure.resourcemanager.cognitiveservices.models.Sku;
import com.azure.resourcemanager.cognitiveservices.models.TrustedCertificateSource;
import com.azure.resourcemanager.cognitiveservices.models.TrustedCertificateSourceType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Accounts Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/CreateAccountWithTrustedCertificateSources.json
     */
    /**
     * Sample code: Create an account with private CA trust.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void createAnAccountWithPrivateCATrust(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accounts().define("myAccount").withExistingResourceGroup("myResourceGroup").withRegion("eastus")
            .withProperties(new AccountProperties().withCustomSubDomainName("myAccount")
                .withAllowProjectManagement(true)
                .withTrustedCertificateSources(Arrays.asList(new TrustedCertificateSource()
                    .withType(TrustedCertificateSourceType.AZURE_KEY_VAULT)
                    .withResourceId(
                        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myKeyVault")
                    .withCertificateBundleReferences(Arrays.asList(
                        new CertificateBundleReference().withSecretName("fakeTokenPlaceholder")
                            .withSecretVersion("fakeTokenPlaceholder"),
                        new CertificateBundleReference().withSecretName("fakeTokenPlaceholder")
                            .withSecretVersion("fakeTokenPlaceholder"))))))
            .withKind("AIServices").withSku(new Sku().withName("S0"))
            .withIdentity(new Identity().withType(ResourceIdentityType.SYSTEM_ASSIGNED)).create();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
