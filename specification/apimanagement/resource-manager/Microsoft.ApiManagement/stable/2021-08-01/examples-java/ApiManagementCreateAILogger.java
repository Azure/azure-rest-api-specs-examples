import com.azure.resourcemanager.apimanagement.models.LoggerType;
import java.util.HashMap;
import java.util.Map;

/** Samples for Logger CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateAILogger.json
     */
    /**
     * Sample code: ApiManagementCreateAILogger.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateAILogger(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .loggers()
            .define("loggerId")
            .withExistingService("rg1", "apimService1")
            .withLoggerType(LoggerType.APPLICATION_INSIGHTS)
            .withDescription("adding a new logger")
            .withCredentials(mapOf("instrumentationKey", "11................a1"))
            .create();
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
