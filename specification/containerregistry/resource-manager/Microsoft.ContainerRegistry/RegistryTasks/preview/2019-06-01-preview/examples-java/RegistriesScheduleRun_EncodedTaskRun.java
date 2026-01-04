
import com.azure.resourcemanager.containerregistry.models.AgentProperties;
import com.azure.resourcemanager.containerregistry.models.EncodedTaskRunRequest;
import com.azure.resourcemanager.containerregistry.models.OS;
import com.azure.resourcemanager.containerregistry.models.PlatformProperties;
import com.azure.resourcemanager.containerregistry.models.SetValue;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Registries ScheduleRun.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/RegistryTasks/preview/2019-06-01-
     * preview/examples/RegistriesScheduleRun_EncodedTaskRun.json
     */
    /**
     * Sample code: Registries_ScheduleRun_EncodedTaskRun.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void registriesScheduleRunEncodedTaskRun(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getRegistries()
            .scheduleRun("myResourceGroup", "myRegistry", new EncodedTaskRunRequest()
                .withEncodedTaskContent("fakeTokenPlaceholder").withEncodedValuesContent("fakeTokenPlaceholder")
                .withValues(Arrays.asList(
                    new SetValue().withName("mytestargument").withValue("mytestvalue").withIsSecret(false),
                    new SetValue().withName("mysecrettestargument").withValue("mysecrettestvalue").withIsSecret(true)))
                .withPlatform(new PlatformProperties().withOs(OS.LINUX))
                .withAgentConfiguration(new AgentProperties().withCpu(2)), com.azure.core.util.Context.NONE);
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
