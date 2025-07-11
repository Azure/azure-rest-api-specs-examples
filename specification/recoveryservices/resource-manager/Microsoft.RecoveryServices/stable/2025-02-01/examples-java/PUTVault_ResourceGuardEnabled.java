
import com.azure.resourcemanager.recoveryservices.models.CmkKekIdentity;
import com.azure.resourcemanager.recoveryservices.models.CmkKeyVaultProperties;
import com.azure.resourcemanager.recoveryservices.models.IdentityData;
import com.azure.resourcemanager.recoveryservices.models.InfrastructureEncryptionState;
import com.azure.resourcemanager.recoveryservices.models.PublicNetworkAccess;
import com.azure.resourcemanager.recoveryservices.models.ResourceIdentityType;
import com.azure.resourcemanager.recoveryservices.models.Sku;
import com.azure.resourcemanager.recoveryservices.models.SkuName;
import com.azure.resourcemanager.recoveryservices.models.UserIdentity;
import com.azure.resourcemanager.recoveryservices.models.VaultProperties;
import com.azure.resourcemanager.recoveryservices.models.VaultPropertiesEncryption;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Vaults CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/
     * PUTVault_ResourceGuardEnabled.json
     */
    /**
     * Sample code: Create or Update Vault performing critical operation With MUA.
     * 
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void createOrUpdateVaultPerformingCriticalOperationWithMUA(
        com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager.vaults().define("swaggerExample").withRegion("West US")
            .withExistingResourceGroup("Default-RecoveryServices-ResourceGroup")
            .withIdentity(new IdentityData().withType(ResourceIdentityType.USER_ASSIGNED)
                .withUserAssignedIdentities(mapOf(
                    "/subscriptions/85bf5e8c-3084-4f42-add2-746ebb7e97b2/resourcegroups/defaultrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/examplemsi",
                    new UserIdentity())))
            .withProperties(new VaultProperties().withEncryption(new VaultPropertiesEncryption()
                .withKeyVaultProperties(new CmkKeyVaultProperties().withKeyUri("fakeTokenPlaceholder"))
                .withKekIdentity(new CmkKekIdentity().withUserAssignedIdentity(
                    "/subscriptions/85bf5e8c-3084-4f42-add2-746ebb7e97b2/resourcegroups/defaultrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/examplemsi"))
                .withInfrastructureEncryption(InfrastructureEncryptionState.ENABLED))
                .withPublicNetworkAccess(PublicNetworkAccess.ENABLED)
                .withResourceGuardOperationRequests(Arrays.asList(
                    "/subscriptions/38304e13-357e-405e-9e9a-220351dcce8c/resourcegroups/ankurResourceGuard1/providers/Microsoft.DataProtection/resourceGuards/ResourceGuard38-1/modifyEncryptionSettings/default")))
            .withSku(new Sku().withName(SkuName.STANDARD)).create();
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
