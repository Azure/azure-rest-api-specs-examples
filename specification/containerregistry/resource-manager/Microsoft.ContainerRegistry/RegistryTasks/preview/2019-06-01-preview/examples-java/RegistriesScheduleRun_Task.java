
import com.azure.resourcemanager.containerregistry.models.Argument;
import com.azure.resourcemanager.containerregistry.models.OverrideTaskStepProperties;
import com.azure.resourcemanager.containerregistry.models.SetValue;
import com.azure.resourcemanager.containerregistry.models.TaskRunRequest;
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
     * preview/examples/RegistriesScheduleRun_Task.json
     */
    /**
     * Sample code: Registries_ScheduleRun_Task.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void registriesScheduleRunTask(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getRegistries().scheduleRun("myResourceGroup",
            "myRegistry",
            new TaskRunRequest().withTaskId("myTask").withOverrideTaskStepProperties(new OverrideTaskStepProperties()
                .withFile("overriddenDockerfile")
                .withArguments(Arrays.asList(
                    new Argument().withName("mytestargument").withValue("mytestvalue").withIsSecret(false),
                    new Argument().withName("mysecrettestargument").withValue("mysecrettestvalue").withIsSecret(true)))
                .withTarget("build")
                .withValues(
                    Arrays.asList(new SetValue().withName("mytestname").withValue("mytestvalue").withIsSecret(false),
                        new SetValue().withName("mysecrettestname").withValue("mysecrettestvalue").withIsSecret(true)))
                .withUpdateTriggerToken("fakeTokenPlaceholder")),
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
