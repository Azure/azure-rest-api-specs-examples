
import com.azure.resourcemanager.containerservice.fluent.models.AgentPoolInner;
import com.azure.resourcemanager.containerservice.models.AgentPoolMode;
import com.azure.resourcemanager.containerservice.models.GpuDriver;
import com.azure.resourcemanager.containerservice.models.GpuProfile;
import com.azure.resourcemanager.containerservice.models.ManagementMode;
import com.azure.resourcemanager.containerservice.models.NvidiaGPUProfile;
import com.azure.resourcemanager.containerservice.models.OSType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for AgentPools CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01/AgentPoolsCreate_ManagedGPU.json
     */
    /**
     * Sample code: Create Agent Pool with fully managed GPU experience.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void createAgentPoolWithFullyManagedGPUExperience(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getAgentPools().createOrUpdate("rg1", "clustername1", "agentpool1",
            new AgentPoolInner().withCount(3).withVmSize("Standard_NC4as_T4_v3").withOsType(OSType.LINUX)
                .withMode(AgentPoolMode.USER)
                .withGpuProfile(new GpuProfile().withDriver(GpuDriver.INSTALL)
                    .withNvidia(new NvidiaGPUProfile().withManagementMode(ManagementMode.MANAGED))),
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
