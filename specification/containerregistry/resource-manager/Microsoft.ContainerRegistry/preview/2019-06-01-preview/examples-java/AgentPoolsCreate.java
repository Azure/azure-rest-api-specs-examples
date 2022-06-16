import com.azure.core.util.Context;
import com.azure.resourcemanager.containerregistry.fluent.models.AgentPoolInner;
import com.azure.resourcemanager.containerregistry.models.OS;
import java.util.HashMap;
import java.util.Map;

/** Samples for AgentPools Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/AgentPoolsCreate.json
     */
    /**
     * Sample code: AgentPools_Create.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void agentPoolsCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getAgentPools()
            .create(
                "myResourceGroup",
                "myRegistry",
                "myAgentPool",
                new AgentPoolInner()
                    .withLocation("WESTUS")
                    .withTags(mapOf("key", "value"))
                    .withCount(1)
                    .withTier("S1")
                    .withOs(OS.LINUX),
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
