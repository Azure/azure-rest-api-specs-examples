
import com.azure.resourcemanager.purestorageblock.models.ManagedServiceIdentity;
import com.azure.resourcemanager.purestorageblock.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.purestorageblock.models.StoragePoolProperties;
import com.azure.resourcemanager.purestorageblock.models.UserAssignedIdentity;
import com.azure.resourcemanager.purestorageblock.models.VnetInjection;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for StoragePools Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/StoragePools_Create_MaximumSet_Gen.json
     */
    /**
     * Sample code: StoragePools_Create.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void storagePoolsCreate(com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.storagePools().define("storagePoolname").withRegion("lonlc").withExistingResourceGroup("rgpurestorage")
            .withTags(mapOf("key7593", "fakeTokenPlaceholder"))
            .withProperties(new StoragePoolProperties().withAvailabilityZone("vknyl")
                .withVnetInjection(new VnetInjection().withSubnetId("tnlctolrxdvnkjiphlrdxq")
                    .withVnetId("zbumtytyqwewjcyckwqchiypshv"))
                .withProvisionedBandwidthMbPerSec(17L).withReservationResourceId("xiowoxnbtcotutcmmrofvgdi"))
            .withIdentity(new ManagedServiceIdentity().withType(ManagedServiceIdentityType.NONE)
                .withUserAssignedIdentities(mapOf("key4211", new UserAssignedIdentity())))
            .create();
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
