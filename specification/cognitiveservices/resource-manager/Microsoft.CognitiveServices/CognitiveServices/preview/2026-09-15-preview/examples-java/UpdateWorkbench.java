
import com.azure.resourcemanager.cognitiveservices.models.Workbench;
import com.azure.resourcemanager.cognitiveservices.models.WorkbenchUpdateProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Workbenches Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-15-preview/UpdateWorkbench.json
     */
    /**
     * Sample code: UpdateWorkbench.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void updateWorkbench(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        Workbench resource = manager.workbenches().getWithResponse("rgcognitiveservices", "myAccount", "myProject",
            "myWorkbench", com.azure.core.util.Context.NONE).getValue();
        resource.update().withProperties(new WorkbenchUpdateProperties().withIdleTimeBeforeShutdown("PT1H")).apply();
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
