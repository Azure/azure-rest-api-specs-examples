import com.azure.resourcemanager.recoveryservices.models.IdentityData;
import com.azure.resourcemanager.recoveryservices.models.ResourceIdentityType;
import com.azure.resourcemanager.recoveryservices.models.UserIdentity;
import com.azure.resourcemanager.recoveryservices.models.Vault;
import java.util.HashMap;
import java.util.Map;

/** Samples for Vaults Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2023-02-01/examples/PATCHVault_WithUserAssignedIdentity.json
     */
    /**
     * Sample code: Update Resource With User Assigned Identity.
     *
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void updateResourceWithUserAssignedIdentity(
        com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        Vault resource =
            manager
                .vaults()
                .getByResourceGroupWithResponse("HelloWorld", "swaggerExample", com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withTags(mapOf("PatchKey", "PatchKeyUpdated"))
            .withIdentity(
                new IdentityData()
                    .withType(ResourceIdentityType.USER_ASSIGNED)
                    .withUserAssignedIdentities(
                        mapOf(
                            "/subscriptions/85bf5e8c-3084-4f42-add2-746ebb7e97b2/resourcegroups/defaultrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/examplemsi",
                            new UserIdentity())))
            .apply();
    }

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
