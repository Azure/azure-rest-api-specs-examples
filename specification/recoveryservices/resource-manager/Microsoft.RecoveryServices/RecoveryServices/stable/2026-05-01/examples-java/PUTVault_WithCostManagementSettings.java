
import com.azure.resourcemanager.recoveryservices.models.CostManagementSettings;
import com.azure.resourcemanager.recoveryservices.models.GranularityLevel;
import com.azure.resourcemanager.recoveryservices.models.IdentityData;
import com.azure.resourcemanager.recoveryservices.models.PublicNetworkAccess;
import com.azure.resourcemanager.recoveryservices.models.ResourceIdentityType;
import com.azure.resourcemanager.recoveryservices.models.Sku;
import com.azure.resourcemanager.recoveryservices.models.SkuName;
import com.azure.resourcemanager.recoveryservices.models.VaultProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Vaults CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/PUTVault_WithCostManagementSettings.json
     */
    /**
     * Sample code: Create or Update Vault With Cost Management Settings.
     * 
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void createOrUpdateVaultWithCostManagementSettings(
        com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager.vaults().define("swaggerExample").withRegion("West US")
            .withExistingResourceGroup("Default-RecoveryServices-ResourceGroup")
            .withProperties(
                new VaultProperties().withPublicNetworkAccess(PublicNetworkAccess.ENABLED).withCostManagementSettings(
                    new CostManagementSettings().withGranularityLevel(GranularityLevel.PROTECTED_ITEM_LEVEL)))
            .withIdentity(new IdentityData().withType(ResourceIdentityType.SYSTEM_ASSIGNED))
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
