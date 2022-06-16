import com.azure.core.util.Context;
import com.azure.resourcemanager.containerservice.fluent.models.AgentPoolInner;
import com.azure.resourcemanager.containerservice.models.AgentPoolMode;
import com.azure.resourcemanager.containerservice.models.OSType;
import com.azure.resourcemanager.containerservice.models.WorkloadRuntime;
import java.util.HashMap;
import java.util.Map;

/** Samples for AgentPools CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-10-01/examples/AgentPoolsCreate_WasmWasi.json
     */
    /**
     * Sample code: Create Agent Pool with Krustlet and the WASI runtime.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAgentPoolWithKrustletAndTheWASIRuntime(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .kubernetesClusters()
            .manager()
            .serviceClient()
            .getAgentPools()
            .createOrUpdate(
                "rg1",
                "clustername1",
                "agentpool1",
                new AgentPoolInner()
                    .withCount(3)
                    .withVmSize("Standard_DS2_v2")
                    .withOsDiskSizeGB(64)
                    .withWorkloadRuntime(WorkloadRuntime.WASM_WASI)
                    .withOsType(OSType.LINUX)
                    .withMode(AgentPoolMode.USER)
                    .withOrchestratorVersion(""),
                Context.NONE);
    }

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
