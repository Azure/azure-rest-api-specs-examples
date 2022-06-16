import com.azure.core.util.Context;
import java.util.HashMap;
import java.util.Map;

/** Samples for Configurations Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/ChangeHttpConnectivityDisable.json
     */
    /**
     * Sample code: Disable HTTP connectivity.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void disableHTTPConnectivity(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager
            .configurations()
            .update("rg1", "cluster1", "gateway", mapOf("restAuthCredential.isEnabled", "false"), Context.NONE);
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
