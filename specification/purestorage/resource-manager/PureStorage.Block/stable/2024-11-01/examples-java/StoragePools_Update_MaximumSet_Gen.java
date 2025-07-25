
import com.azure.resourcemanager.purestorageblock.models.ManagedServiceIdentity;
import com.azure.resourcemanager.purestorageblock.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.purestorageblock.models.StoragePool;
import com.azure.resourcemanager.purestorageblock.models.StoragePoolUpdateProperties;
import com.azure.resourcemanager.purestorageblock.models.UserAssignedIdentity;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for StoragePools Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/StoragePools_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: StoragePools_Update.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void storagePoolsUpdate(com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        StoragePool resource = manager.storagePools()
            .getByResourceGroupWithResponse("rgpurestorage", "storagePoolname", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("key9065", "fakeTokenPlaceholder"))
            .withIdentity(new ManagedServiceIdentity().withType(ManagedServiceIdentityType.NONE)
                .withUserAssignedIdentities(mapOf("key4211", new UserAssignedIdentity())))
            .withProperties(new StoragePoolUpdateProperties().withProvisionedBandwidthMbPerSec(23L)).apply();
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
