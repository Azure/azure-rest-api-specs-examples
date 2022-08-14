import com.azure.core.util.Context;
import com.azure.resourcemanager.automation.models.Module;
import java.util.HashMap;
import java.util.Map;

/** Samples for Python2Package Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/updatePython2Package.json
     */
    /**
     * Sample code: Update a module.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void updateAModule(com.azure.resourcemanager.automation.AutomationManager manager) {
        Module resource =
            manager
                .python2Packages()
                .getWithResponse("rg", "MyAutomationAccount", "MyPython2Package", Context.NONE)
                .getValue();
        resource.update().withTags(mapOf()).apply();
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
