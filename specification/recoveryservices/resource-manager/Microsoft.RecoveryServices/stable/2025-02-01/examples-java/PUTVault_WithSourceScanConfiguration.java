
import com.azure.resourcemanager.recoveryservices.models.AssociatedIdentity;
import com.azure.resourcemanager.recoveryservices.models.IdentityData;
import com.azure.resourcemanager.recoveryservices.models.IdentityType;
import com.azure.resourcemanager.recoveryservices.models.PublicNetworkAccess;
import com.azure.resourcemanager.recoveryservices.models.ResourceIdentityType;
import com.azure.resourcemanager.recoveryservices.models.SecuritySettings;
import com.azure.resourcemanager.recoveryservices.models.Sku;
import com.azure.resourcemanager.recoveryservices.models.SkuName;
import com.azure.resourcemanager.recoveryservices.models.SourceScanConfiguration;
import com.azure.resourcemanager.recoveryservices.models.State;
import com.azure.resourcemanager.recoveryservices.models.VaultProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Vaults CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/
     * PUTVault_WithSourceScanConfiguration.json
     */
    /**
     * Sample code: Create or Update Vault with Source scan configuration.
     * 
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void createOrUpdateVaultWithSourceScanConfiguration(
        com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager.vaults().define("swaggerExample").withRegion("West US")
            .withExistingResourceGroup("Default-RecoveryServices-ResourceGroup")
            .withIdentity(new IdentityData().withType(ResourceIdentityType.SYSTEM_ASSIGNED))
            .withProperties(new VaultProperties().withPublicNetworkAccess(PublicNetworkAccess.ENABLED)
                .withSecuritySettings(new SecuritySettings().withSourceScanConfiguration(
                    new SourceScanConfiguration().withState(State.ENABLED).withSourceScanIdentity(
                        new AssociatedIdentity().withOperationIdentityType(IdentityType.SYSTEM_ASSIGNED)))))
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
