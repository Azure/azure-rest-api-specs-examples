
import com.azure.resourcemanager.automation.models.DscConfiguration;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for DscConfiguration Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/updateDscConfiguration.
     * json
     */
    /**
     * Sample code: Create or Update Configuration.
     * 
     * @param manager Entry point to AutomationManager.
     */
    public static void createOrUpdateConfiguration(com.azure.resourcemanager.automation.AutomationManager manager) {
        DscConfiguration resource = manager.dscConfigurations()
            .getWithResponse("rg", "myAutomationAccount18", "SetupServer", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("Hello", "World")).withName("SetupServer").apply();
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
