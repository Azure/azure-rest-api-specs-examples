
import com.azure.resourcemanager.cognitiveservices.models.Account;
import com.azure.resourcemanager.cognitiveservices.models.AccountProperties;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Accounts Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/UpdateAccountClearTrustedCertificateSources.json
     */
    /**
     * Sample code: Clear account private CA trust using an empty array.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void clearAccountPrivateCATrustUsingAnEmptyArray(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        Account resource = manager.accounts()
            .getByResourceGroupWithResponse("myResourceGroup", "myAccount", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withProperties(new AccountProperties().withTrustedCertificateSources(Arrays.asList()))
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
