
import com.azure.resourcemanager.datadog.models.DatadogMonitorResource;
import com.azure.resourcemanager.datadog.models.MonitorUpdateProperties;
import com.azure.resourcemanager.datadog.models.MonitoringStatus;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Monitors Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/Monitors_Update.json
     */
    /**
     * Sample code: Monitors_Update.
     * 
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void monitorsUpdate(com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        DatadogMonitorResource resource = manager.monitors()
            .getByResourceGroupWithResponse("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("Environment", "Dev"))
            .withProperties(new MonitorUpdateProperties().withMonitoringStatus(MonitoringStatus.ENABLED)).apply();
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
