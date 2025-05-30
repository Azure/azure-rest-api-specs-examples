
import com.azure.resourcemanager.appcontainers.models.BuilderResource;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Builders Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/Builders_Update.json
     */
    /**
     * Sample code: Builders_Update_0.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void buildersUpdate0(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        BuilderResource resource = manager.builders()
            .getByResourceGroupWithResponse("rg", "testBuilder", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("mytag1", "myvalue1")).apply();
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
