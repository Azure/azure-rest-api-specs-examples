import com.azure.core.util.Context;
import com.azure.resourcemanager.cdn.models.AfdEndpointUpdateParameters;
import com.azure.resourcemanager.cdn.models.EnabledState;
import java.util.HashMap;
import java.util.Map;

/** Samples for AfdEndpoints Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDEndpoints_Update.json
     */
    /**
     * Sample code: AFDEndpoints_Update.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void aFDEndpointsUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getAfdEndpoints()
            .update(
                "RG",
                "profile1",
                "endpoint1",
                new AfdEndpointUpdateParameters().withTags(mapOf()).withEnabledState(EnabledState.ENABLED),
                Context.NONE);
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
