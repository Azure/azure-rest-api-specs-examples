
import com.azure.resourcemanager.containerservice.fluent.models.AgentPoolInner;
import com.azure.resourcemanager.containerservice.models.AgentPoolType;
import com.azure.resourcemanager.containerservice.models.AutoScaleProfile;
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
     * x-ms-original-file: 2026-05-01/AgentPoolsCreate_TypeVirtualMachines_Autoscale.json
     */
    /**
     * Sample code: Create Agent Pool with VirtualMachines pool type with autoscaling enabled.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void createAgentPoolWithVirtualMachinesPoolTypeWithAutoscalingEnabled(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getAgentPools().createOrUpdate("rg1", "clustername1", "agentpool1",
            new AgentPoolInner().withOsType(OSType.LINUX).withTypePropertiesType(AgentPoolType.VIRTUAL_MACHINES)
                .withOrchestratorVersion("1.29.0").withTags(mapOf("name1", "val1"))
                .withNodeLabels(mapOf("key1", "fakeTokenPlaceholder"))
                .withNodeTaints(Arrays.asList("Key1=Value1:NoSchedule"))
                .withVirtualMachinesProfile(new VirtualMachinesProfile().withScale(new ScaleProfile().withAutoscale(
                    Arrays.asList(new AutoScaleProfile().withSize("Standard_D2_v2").withMinCount(1).withMaxCount(5))))),
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
