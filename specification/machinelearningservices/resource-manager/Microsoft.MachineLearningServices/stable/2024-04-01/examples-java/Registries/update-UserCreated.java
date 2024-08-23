
import com.azure.resourcemanager.machinelearning.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.machinelearning.models.PartialSku;
import com.azure.resourcemanager.machinelearning.models.Registry;
import com.azure.resourcemanager.machinelearning.models.RegistryPartialManagedServiceIdentity;
import com.azure.resourcemanager.machinelearning.models.SkuTier;
import com.azure.resourcemanager.machinelearning.models.UserAssignedIdentity;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Registries Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Registries/update-UserCreated.json
     */
    /**
     * Sample code: Update Registry with user created accounts.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void updateRegistryWithUserCreatedAccounts(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        Registry resource = manager.registries()
            .getByResourceGroupWithResponse("test-rg", "string", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf())
            .withIdentity(new RegistryPartialManagedServiceIdentity().withType(ManagedServiceIdentityType.USER_ASSIGNED)
                .withUserAssignedIdentities(mapOf("string", new UserAssignedIdentity())))
            .withSku(new PartialSku().withName("string").withTier(SkuTier.BASIC).withSize("string").withFamily("string")
                .withCapacity(1))
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
