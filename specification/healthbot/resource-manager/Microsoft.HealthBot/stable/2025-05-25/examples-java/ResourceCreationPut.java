
import com.azure.resourcemanager.healthbot.models.Identity;
import com.azure.resourcemanager.healthbot.models.ResourceIdentityType;
import com.azure.resourcemanager.healthbot.models.Sku;
import com.azure.resourcemanager.healthbot.models.SkuName;
import com.azure.resourcemanager.healthbot.models.UserAssignedIdentity;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Bots Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthbot/resource-manager/Microsoft.HealthBot/stable/2025-05-25/examples/ResourceCreationPut.json
     */
    /**
     * Sample code: BotCreate.
     * 
     * @param manager Entry point to HealthbotManager.
     */
    public static void botCreate(com.azure.resourcemanager.healthbot.HealthbotManager manager) {
        manager.bots().define("samplebotname").withRegion("East US").withExistingResourceGroup("healthbotClient")
            .withSku(new Sku().withName(SkuName.F0))
            .withIdentity(new Identity().withType(ResourceIdentityType.SYSTEM_ASSIGNED_USER_ASSIGNED)
                .withUserAssignedIdentities(mapOf(
                    "/subscriptions/subscription-id/resourcegroups/myrg/providers/microsoft.managedidentity/userassignedidentities/my-mi",
                    new UserAssignedIdentity(),
                    "/subscriptions/subscription-id/resourcegroups/myrg/providers/microsoft.managedidentity/userassignedidentities/my-mi2",
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
