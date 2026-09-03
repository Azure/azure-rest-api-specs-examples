
import com.azure.resourcemanager.cognitiveservices.models.AccountProperties;
import com.azure.resourcemanager.cognitiveservices.models.Identity;
import com.azure.resourcemanager.cognitiveservices.models.ResourceIdentityType;
import com.azure.resourcemanager.cognitiveservices.models.Sku;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Accounts Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-15-preview/CreateAccountMin.json
     */
    /**
     * Sample code: Create Account Min.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void createAccountMin(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accounts().define("testCreate1").withExistingResourceGroup("myResourceGroup").withRegion("West US")
            .withProperties(new AccountProperties()).withKind("CognitiveServices").withSku(new Sku().withName("S0"))
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
