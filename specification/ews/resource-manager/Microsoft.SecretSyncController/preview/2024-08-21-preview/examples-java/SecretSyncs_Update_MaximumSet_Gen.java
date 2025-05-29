
import com.azure.resourcemanager.secretsstoreextension.models.KubernetesSecretObjectMapping;
import com.azure.resourcemanager.secretsstoreextension.models.KubernetesSecretType;
import com.azure.resourcemanager.secretsstoreextension.models.SecretSync;
import com.azure.resourcemanager.secretsstoreextension.models.SecretSyncUpdateProperties;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for SecretSyncs Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-08-21-preview/SecretSyncs_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: SecretSyncs_Update.
     * 
     * @param manager Entry point to SecretsStoreExtensionManager.
     */
    public static void
        secretSyncsUpdate(com.azure.resourcemanager.secretsstoreextension.SecretsStoreExtensionManager manager) {
        SecretSync resource = manager.secretSyncs().getByResourceGroupWithResponse("rg-ssc-example",
            "secretsync-ssc-example", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("example-tag", "example-tag-value")).withProperties(
            new SecretSyncUpdateProperties().withSecretProviderClassName("fakeTokenPlaceholder").withServiceAccountName(
                "fcldqfdfpktndlntuoxicsftelhefevovmlycflfwzckvamiqjnjugandqaqqeccsbzztfmmeunvhsafgerbcsdbnmsyqivygornebbkusuvphwghgouxvcbvmbydqjzoxextnyowsnyymadniwdrrxtogeveldpejixmsrzzfqkquaxdpzwvecevqwasxgxxchrfa")
                .withKubernetesSecretType(KubernetesSecretType.OPAQUE).withForceSynchronization("arbitrarystring")
                .withObjectSecretMapping(Arrays.asList(new KubernetesSecretObjectMapping().withSourcePath(
                    "ssrzmbvdiomkvzrdsyilwlfzicfydnbjwjsnohrppkukjddrunfslkrnexunuckmghixdssposvndpiqchpqrkjuqbapoisvqdvgstvdonsmlpsmticfvuhqlofpaxfdg")
                    .withTargetKey("fakeTokenPlaceholder"))))
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
