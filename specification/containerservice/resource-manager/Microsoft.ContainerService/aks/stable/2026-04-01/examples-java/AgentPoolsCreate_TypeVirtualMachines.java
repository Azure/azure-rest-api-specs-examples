
import com.azure.resourcemanager.containerservice.fluent.models.AgentPoolInner;
import com.azure.resourcemanager.containerservice.models.AgentPoolType;
import com.azure.resourcemanager.containerservice.models.ManualScaleProfile;
import com.azure.resourcemanager.containerservice.models.OSType;
import com.azure.resourcemanager.containerservice.models.ScaleProfile;
import com.azure.resourcemanager.containerservice.models.VirtualMachinesProfile;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for AgentPools CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/AgentPoolsCreate_TypeVirtualMachines.json
     */
    /**
     * Sample code: Create Agent Pool with VirtualMachines pool type.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void createAgentPoolWithVirtualMachinesPoolType(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getAgentPools().createOrUpdate("rg1", "clustername1", "agentpool1",
            new AgentPoolInner().withOsType(OSType.LINUX).withTypePropertiesType(AgentPoolType.VIRTUAL_MACHINES)
                .withOrchestratorVersion("1.9.6").withTags(mapOf("name1", "val1"))
                .withNodeLabels(mapOf("key1", "fakeTokenPlaceholder"))
                .withNodeTaints(Arrays.asList("Key1=Value1:NoSchedule"))
                .withVirtualMachinesProfile(new VirtualMachinesProfile().withScale(new ScaleProfile()
                    .withManual(Arrays.asList(new ManualScaleProfile().withSize("Standard_D2_v2").withCount(3),
                        new ManualScaleProfile().withSize("Standard_D2_v3").withCount(2))))),
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
