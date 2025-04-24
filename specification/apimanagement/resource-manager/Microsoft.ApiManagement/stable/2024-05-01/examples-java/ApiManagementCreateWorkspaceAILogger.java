
import com.azure.resourcemanager.apimanagement.fluent.models.LoggerContractInner;
import com.azure.resourcemanager.apimanagement.models.LoggerType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for WorkspaceLogger CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateWorkspaceAILogger.json
     */
    /**
     * Sample code: ApiManagementCreateWorkspaceAILogger.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementCreateWorkspaceAILogger(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceLoggers().createOrUpdateWithResponse("rg1", "apimService1", "wks1", "loggerId",
            new LoggerContractInner().withLoggerType(LoggerType.APPLICATION_INSIGHTS)
                .withDescription("adding a new logger").withCredentials(
                    mapOf("instrumentationKey", "fakeTokenPlaceholder")),
            null, com.azure.core.util.Context.NONE);
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
