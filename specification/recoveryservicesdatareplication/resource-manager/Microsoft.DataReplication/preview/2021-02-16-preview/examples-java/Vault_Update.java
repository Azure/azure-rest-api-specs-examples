
import com.azure.resourcemanager.recoveryservicesdatareplication.models.ReplicationVaultType;
import com.azure.resourcemanager.recoveryservicesdatareplication.models.VaultModel;
import com.azure.resourcemanager.recoveryservicesdatareplication.models.VaultModelProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Vault Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-
     * preview/examples/Vault_Update.json
     */
    /**
     * Sample code: Vault_Update.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void vaultUpdate(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        VaultModel resource = manager.vaults()
            .getByResourceGroupWithResponse("rgrecoveryservicesdatareplication", "4", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("key8872", "fakeTokenPlaceholder"))
            .withProperties(new VaultModelProperties().withVaultType(ReplicationVaultType.DISASTER_RECOVERY)).apply();
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
