
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Configurations Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/
     * ChangeHttpConnectivityEnable.json
     */
    /**
     * Sample code: Enable HTTP connectivity.
     * 
     * @param manager Entry point to HDInsightManager.
     */
    public static void enableHTTPConnectivity(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.configurations().update("rg1", "cluster1", "gateway",
            mapOf("restAuthCredential.isEnabled", "fakeTokenPlaceholder", "restAuthCredential.password",
                "fakeTokenPlaceholder", "restAuthCredential.username", "fakeTokenPlaceholder"),
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
