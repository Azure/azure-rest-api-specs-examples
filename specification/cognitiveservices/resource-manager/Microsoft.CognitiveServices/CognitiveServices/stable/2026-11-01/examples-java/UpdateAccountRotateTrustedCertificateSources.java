
import com.azure.resourcemanager.cognitiveservices.models.Account;
import com.azure.resourcemanager.cognitiveservices.models.AccountProperties;
import com.azure.resourcemanager.cognitiveservices.models.CertificateBundleReference;
import com.azure.resourcemanager.cognitiveservices.models.TrustedCertificateSource;
import com.azure.resourcemanager.cognitiveservices.models.TrustedCertificateSourceType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Accounts Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/UpdateAccountRotateTrustedCertificateSources.json
     */
    /**
     * Sample code: Rotate a pinned CA bundle while retaining other references.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void rotateAPinnedCABundleWhileRetainingOtherReferences(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        Account resource = manager.accounts()
            .getByResourceGroupWithResponse("myResourceGroup", "myAccount", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withProperties(new AccountProperties().withTrustedCertificateSources(Arrays.asList(
            new TrustedCertificateSource().withType(TrustedCertificateSourceType.AZURE_KEY_VAULT).withResourceId(
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myKeyVault")
                .withCertificateBundleReferences(Arrays.asList(
                    new CertificateBundleReference().withSecretName("fakeTokenPlaceholder")
                        .withSecretVersion("fakeTokenPlaceholder"),
                    new CertificateBundleReference().withSecretName("fakeTokenPlaceholder")
                        .withSecretVersion("fakeTokenPlaceholder"))))))
            .apply();
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
