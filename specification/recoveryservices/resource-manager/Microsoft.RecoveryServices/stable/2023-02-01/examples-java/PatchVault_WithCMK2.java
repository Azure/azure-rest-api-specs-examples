import com.azure.resourcemanager.recoveryservices.models.CmkKekIdentity;
import com.azure.resourcemanager.recoveryservices.models.IdentityData;
import com.azure.resourcemanager.recoveryservices.models.ResourceIdentityType;
import com.azure.resourcemanager.recoveryservices.models.Vault;
import com.azure.resourcemanager.recoveryservices.models.VaultProperties;
import com.azure.resourcemanager.recoveryservices.models.VaultPropertiesEncryption;
import java.util.HashMap;
import java.util.Map;

/** Samples for Vaults Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2023-02-01/examples/PatchVault_WithCMK2.json
     */
    /**
     * Sample code: Update Resource With CustomerManagedKeys2.
     *
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void updateResourceWithCustomerManagedKeys2(
        com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        Vault resource =
            manager
                .vaults()
                .getByResourceGroupWithResponse("HelloWorld", "swaggerExample", com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withTags(mapOf("PatchKey", "PatchKeyUpdated"))
            .withProperties(
                new VaultProperties()
                    .withEncryption(
                        new VaultPropertiesEncryption()
                            .withKekIdentity(new CmkKekIdentity().withUseSystemAssignedIdentity(true))))
            .withIdentity(new IdentityData().withType(ResourceIdentityType.SYSTEM_ASSIGNED))
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
