
import com.azure.resourcemanager.cognitiveservices.models.AccountProperties;
import com.azure.resourcemanager.cognitiveservices.models.Identity;
import com.azure.resourcemanager.cognitiveservices.models.ResourceIdentityType;
import com.azure.resourcemanager.cognitiveservices.models.Sku;
import com.azure.resourcemanager.cognitiveservices.models.UserAssignedIdentity;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Accounts Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/CreateAccountWithAgentHostingConfiguration.json
     */
    /**
     * Sample code: Create a Foundry account with customer-owned AKS hosting.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void createAFoundryAccountWithCustomerOwnedAKSHosting(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accounts().define("foundryByocAccount").withExistingResourceGroup("myResourceGroup")
            .withRegion("West US").withProperties(new AccountProperties()).withKind("AIServices")
            .withSku(new Sku().withName("S0"))
            .withIdentity(new Identity().withType(ResourceIdentityType.USER_ASSIGNED).withUserAssignedIdentities(mapOf(
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/account-control-plane",
                new UserAssignedIdentity())))
            .create();
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
