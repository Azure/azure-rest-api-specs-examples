```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.containerservice.fluent.models.AgentPoolInner;
import com.azure.resourcemanager.containerservice.models.AgentPoolMode;
import com.azure.resourcemanager.containerservice.models.OSType;
import com.azure.resourcemanager.containerservice.models.ScaleSetEvictionPolicy;
import com.azure.resourcemanager.containerservice.models.ScaleSetPriority;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for AgentPools CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-02-01/examples/AgentPoolsCreate_Update.json
     */
    /**
     * Sample code: Create/Update Agent Pool.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createUpdateAgentPool(com.azure.resourcemanager.AzureResourceManager azure) {
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
                    .withVmSize("Standard_DS1_v2")
                    .withOsType(OSType.LINUX)
                    .withMode(AgentPoolMode.USER)
                    .withOrchestratorVersion("")
                    .withScaleSetPriority(ScaleSetPriority.SPOT)
                    .withScaleSetEvictionPolicy(ScaleSetEvictionPolicy.DELETE)
                    .withTags(mapOf("name1", "val1"))
                    .withNodeLabels(mapOf("key1", "val1"))
                    .withNodeTaints(Arrays.asList("Key1=Value1:NoSchedule")),
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
