
import com.azure.resourcemanager.networkcloud.models.BmcKeySetPrivilegeLevel;
import com.azure.resourcemanager.networkcloud.models.ExtendedLocation;
import com.azure.resourcemanager.networkcloud.models.KeySetUser;
import com.azure.resourcemanager.networkcloud.models.SshPublicKey;
import java.time.OffsetDateTime;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for BmcKeySets CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/BmcKeySets_Create.
     * json
     */
    /**
     * Sample code: Create or update baseboard management controller key set of cluster.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void createOrUpdateBaseboardManagementControllerKeySetOfCluster(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.bmcKeySets().define("bmcKeySetName").withRegion("location")
            .withExistingCluster("resourceGroupName", "clusterName")
            .withExtendedLocation(new ExtendedLocation().withName(
                "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName")
                .withType("CustomLocation"))
            .withAzureGroupId("f110271b-XXXX-4163-9b99-214d91660f0e")
            .withExpiration(OffsetDateTime.parse("2022-12-31T23:59:59.008Z"))
            .withPrivilegeLevel(BmcKeySetPrivilegeLevel.ADMINISTRATOR)
            .withUserList(Arrays.asList(
                new KeySetUser().withAzureUsername("userABC")
                    .withDescription("Needs access for troubleshooting as a part of the support team")
                    .withSshPublicKey(new SshPublicKey().withKeyData("fakeTokenPlaceholder"))
                    .withUserPrincipalName("userABC@contoso.com"),
                new KeySetUser().withAzureUsername("userXYZ")
                    .withDescription("Needs access for troubleshooting as a part of the support team")
                    .withSshPublicKey(new SshPublicKey().withKeyData("fakeTokenPlaceholder"))
                    .withUserPrincipalName("userABC@contoso.com")))
            .withTags(mapOf("key1", "fakeTokenPlaceholder", "key2", "fakeTokenPlaceholder")).create();
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
