
import com.azure.resourcemanager.containerinstance.fluent.models.ContainerGroupInner;
import com.azure.resourcemanager.containerinstance.models.Container;
import com.azure.resourcemanager.containerinstance.models.ContainerGroupPriority;
import com.azure.resourcemanager.containerinstance.models.ContainerGroupRestartPolicy;
import com.azure.resourcemanager.containerinstance.models.ContainerGroupSku;
import com.azure.resourcemanager.containerinstance.models.OperatingSystemTypes;
import com.azure.resourcemanager.containerinstance.models.ResourceRequests;
import com.azure.resourcemanager.containerinstance.models.ResourceRequirements;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for ContainerGroups CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2023-05-01/examples/
     * ContainerGroupsCreatePriority.json
     */
    /**
     * Sample code: ContainerGroupsCreateWithPriority.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void containerGroupsCreateWithPriority(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerGroups().manager().serviceClient().getContainerGroups().createOrUpdate("demo", "demo1",
            new ContainerGroupInner().withLocation("eastus")
                .withContainers(Arrays.asList(new Container().withName("test-container-001").withImage("alpine:latest")
                    .withCommand(Arrays.asList("/bin/sh", "-c", "sleep 10"))
                    .withResources(new ResourceRequirements()
                        .withRequests(new ResourceRequests().withMemoryInGB(1.0).withCpu(1.0)))))
                .withRestartPolicy(ContainerGroupRestartPolicy.NEVER).withOsType(OperatingSystemTypes.LINUX)
                .withSku(ContainerGroupSku.STANDARD).withPriority(ContainerGroupPriority.SPOT),
            com.azure.core.util.Context.NONE);
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
