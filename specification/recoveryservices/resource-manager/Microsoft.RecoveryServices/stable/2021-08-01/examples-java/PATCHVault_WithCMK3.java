import com.azure.core.util.Context;
import com.azure.resourcemanager.recoveryservices.models.CmkKeyVaultProperties;
import com.azure.resourcemanager.recoveryservices.models.IdentityData;
import com.azure.resourcemanager.recoveryservices.models.ResourceIdentityType;
import com.azure.resourcemanager.recoveryservices.models.UserIdentity;
import com.azure.resourcemanager.recoveryservices.models.Vault;
import com.azure.resourcemanager.recoveryservices.models.VaultProperties;
import com.azure.resourcemanager.recoveryservices.models.VaultPropertiesEncryption;
import java.util.HashMap;
import java.util.Map;

/** Samples for Vaults Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2021-08-01/examples/PATCHVault_WithCMK3.json
     */
    /**
     * Sample code: Update Resource With CustomerManagedKeys3.
     *
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void updateResourceWithCustomerManagedKeys3(
        com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        Vault resource =
            manager.vaults().getByResourceGroupWithResponse("HelloWorld", "swaggerExample", Context.NONE).getValue();
        resource
            .update()
            .withTags(mapOf("PatchKey", "PatchKeyUpdated"))
            .withProperties(
                new VaultProperties()
                    .withEncryption(
                        new VaultPropertiesEncryption()
                            .withKeyVaultProperties(
                                new CmkKeyVaultProperties()
                                    .withKeyUri(
                                        "https://cmk2xkv.vault.azure.net/keys/Key1/0767b348bb1a4c07baa6c4ec0055d2b3"))))
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
