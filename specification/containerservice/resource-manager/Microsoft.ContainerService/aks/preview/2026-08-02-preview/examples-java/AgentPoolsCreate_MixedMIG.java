
import com.azure.resourcemanager.containerservice.fluent.models.AgentPoolInner;
import com.azure.resourcemanager.containerservice.models.AgentPoolMode;
import com.azure.resourcemanager.containerservice.models.GpuDriver;
import com.azure.resourcemanager.containerservice.models.GpuInstanceProfile;
import com.azure.resourcemanager.containerservice.models.GpuProfile;
import com.azure.resourcemanager.containerservice.models.ManagementMode;
import com.azure.resourcemanager.containerservice.models.MigStrategy;
import com.azure.resourcemanager.containerservice.models.NvidiaGPUProfile;
import com.azure.resourcemanager.containerservice.models.OSType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for AgentPools CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-02-preview/AgentPoolsCreate_MixedMIG.json
     */
    /**
     * Sample code: Create Agent Pool with mixed MIG profiles.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void createAgentPoolWithMixedMIGProfiles(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getAgentPools().createOrUpdate("rg1", "clustername1", "agentpool1",
            new AgentPoolInner().withCount(3).withVmSize("Standard_ND96asr_v4").withOsType(OSType.LINUX)
                .withMode(AgentPoolMode.USER)
                .withGpuProfile(new GpuProfile().withDriver(GpuDriver.INSTALL)
                    .withNvidia(new NvidiaGPUProfile().withManagementMode(ManagementMode.MANAGED)
                        .withMigStrategy(MigStrategy.MIXED)
                        .withMigProfiles(Arrays.asList(GpuInstanceProfile.MIG3G, GpuInstanceProfile.MIG2G,
                            GpuInstanceProfile.MIG1G, GpuInstanceProfile.MIG1G)))),
            null, null, com.azure.core.util.Context.NONE);
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
