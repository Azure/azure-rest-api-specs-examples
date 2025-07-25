
import com.azure.resourcemanager.networkcloud.models.AgentPool;
import com.azure.resourcemanager.networkcloud.models.AgentPoolUpgradeSettings;
import com.azure.resourcemanager.networkcloud.models.NodePoolAdministratorConfigurationPatch;
import com.azure.resourcemanager.networkcloud.models.SshPublicKey;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for AgentPools Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/AgentPools_Patch.
     * json
     */
    /**
     * Sample code: Patch Kubernetes cluster agent pool.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        patchKubernetesClusterAgentPool(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        AgentPool resource = manager.agentPools().getWithResponse("resourceGroupName", "kubernetesClusterName",
            "agentPoolName", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("key1", "fakeTokenPlaceholder", "key2", "fakeTokenPlaceholder"))
            .withAdministratorConfiguration(new NodePoolAdministratorConfigurationPatch()
                .withSshPublicKeys(Arrays.asList(new SshPublicKey().withKeyData("fakeTokenPlaceholder"))))
            .withCount(3L)
            .withUpgradeSettings(
                new AgentPoolUpgradeSettings().withDrainTimeout(1800L).withMaxSurge("1").withMaxUnavailable("0"))
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
