
import com.azure.resourcemanager.elasticsan.models.EncryptionIdentity;
import com.azure.resourcemanager.elasticsan.models.EncryptionProperties;
import com.azure.resourcemanager.elasticsan.models.EncryptionType;
import com.azure.resourcemanager.elasticsan.models.Identity;
import com.azure.resourcemanager.elasticsan.models.IdentityType;
import com.azure.resourcemanager.elasticsan.models.KeyVaultProperties;
import com.azure.resourcemanager.elasticsan.models.QualityOfService;
import com.azure.resourcemanager.elasticsan.models.StorageTargetType;
import com.azure.resourcemanager.elasticsan.models.UserAssignedIdentity;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for VolumeGroups Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/VolumeGroups_PerformanceCritical_Create_MaximumSet_Gen.json
     */
    /**
     * Sample code: VolumeGroups_PerformanceCritical_Create_MaximumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void volumeGroupsPerformanceCriticalCreateMaximumSetGen(
        com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.volumeGroups().define("volumegroupname").withExistingElasticSan("resourcegroupname", "elasticsanname")
            .withIdentity(new Identity().withType(IdentityType.NONE)
                .withUserAssignedIdentities(mapOf("key2350", new UserAssignedIdentity())))
            .withProtocolType(StorageTargetType.DIRECT_ATTACH)
            .withEncryption(EncryptionType.ENCRYPTION_AT_REST_WITH_PLATFORM_KEY)
            .withEncryptionProperties(new EncryptionProperties()
                .withKeyVaultProperties(new KeyVaultProperties().withKeyName("fakeTokenPlaceholder")
                    .withKeyVersion("fakeTokenPlaceholder").withKeyVaultUri("fakeTokenPlaceholder"))
                .withEncryptionIdentity(new EncryptionIdentity().withEncryptionUserAssignedIdentity("vgbeephfgecgg")))
            .withReservedIops(10000).withReservedMBps(800).withQualityOfService(QualityOfService.PERFORMANCE_CRITICAL)
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
