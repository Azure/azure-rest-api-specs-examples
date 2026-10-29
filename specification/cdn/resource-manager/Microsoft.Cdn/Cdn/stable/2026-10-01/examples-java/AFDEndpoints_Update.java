
import com.azure.resourcemanager.cdn.models.AfdEndpointUpdateParameters;
import com.azure.resourcemanager.cdn.models.EnabledState;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for AfdEndpoints Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/AFDEndpoints_Update.json
     */
    /**
     * Sample code: AFDEndpoints_Update.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void aFDEndpointsUpdate(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getAfdEndpoints().update("RG", "profile1", "endpoint1",
            new AfdEndpointUpdateParameters().withTags(mapOf()).withEnabledState(EnabledState.ENABLED),
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
