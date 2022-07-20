import com.azure.resourcemanager.recoveryservices.models.IdentityData;
import com.azure.resourcemanager.recoveryservices.models.ResourceIdentityType;
import com.azure.resourcemanager.recoveryservices.models.Sku;
import com.azure.resourcemanager.recoveryservices.models.SkuName;
import com.azure.resourcemanager.recoveryservices.models.VaultProperties;
import java.util.HashMap;
import java.util.Map;

/** Samples for Vaults CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2021-08-01/examples/PUTVault.json
     */
    /**
     * Sample code: Create or Update Recovery Services vault.
     *
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void createOrUpdateRecoveryServicesVault(
        com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager
            .vaults()
            .define("swaggerExample")
            .withRegion("West US")
            .withExistingResourceGroup("Default-RecoveryServices-ResourceGroup")
            .withIdentity(new IdentityData().withType(ResourceIdentityType.SYSTEM_ASSIGNED))
            .withProperties(new VaultProperties())
            .withSku(new Sku().withName(SkuName.STANDARD))
            .create();
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
