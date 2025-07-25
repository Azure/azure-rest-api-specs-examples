
import com.azure.resourcemanager.networkcloud.models.BareMetalMachineKeySet;
import com.azure.resourcemanager.networkcloud.models.KeySetUser;
import com.azure.resourcemanager.networkcloud.models.SshPublicKey;
import java.time.OffsetDateTime;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for BareMetalMachineKeySets Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/
     * BareMetalMachineKeySets_Patch.json
     */
    /**
     * Sample code: Patch bare metal machine key set of cluster.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        patchBareMetalMachineKeySetOfCluster(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        BareMetalMachineKeySet resource = manager.bareMetalMachineKeySets().getWithResponse("resourceGroupName",
            "clusterName", "bareMetalMachineKeySetName", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("key1", "fakeTokenPlaceholder", "key2", "fakeTokenPlaceholder"))
            .withExpiration(OffsetDateTime.parse("2022-12-31T23:59:59.008Z"))
            .withJumpHostsAllowed(Arrays.asList("192.0.2.1", "192.0.2.5"))
            .withUserList(Arrays.asList(
                new KeySetUser().withAzureUsername("userABC")
                    .withDescription("Needs access for troubleshooting as a part of the support team")
                    .withSshPublicKey(new SshPublicKey().withKeyData("fakeTokenPlaceholder"))
                    .withUserPrincipalName("userABC@contoso.com"),
                new KeySetUser().withAzureUsername("userXYZ")
                    .withDescription("Needs access for troubleshooting as a part of the support team")
                    .withSshPublicKey(new SshPublicKey().withKeyData("fakeTokenPlaceholder"))
                    .withUserPrincipalName("userABC@contoso.com")))
            .apply();
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
