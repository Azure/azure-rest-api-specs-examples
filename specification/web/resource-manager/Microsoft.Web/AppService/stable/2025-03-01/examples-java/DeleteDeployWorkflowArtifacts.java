
import com.azure.resourcemanager.appservice.models.WorkflowArtifacts;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for WebApps DeployWorkflowArtifacts.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * DeleteDeployWorkflowArtifacts.json
     */
    /**
     * Sample code: Delete workflow artifacts.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteWorkflowArtifacts(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().deployWorkflowArtifactsWithResponse("testrg123",
            "testsite2", new WorkflowArtifacts().withFilesToDelete(Arrays.asList("test/workflow.json", "test/")),
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
