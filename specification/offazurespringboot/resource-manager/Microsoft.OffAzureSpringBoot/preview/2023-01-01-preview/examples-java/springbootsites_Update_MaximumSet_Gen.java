
import com.azure.resourcemanager.springappdiscovery.models.SpringbootsitesModel;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Springbootsites Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/
     * examples/springbootsites_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: springbootsites_Update_MaximumSet_Gen.
     * 
     * @param manager Entry point to SpringAppDiscoveryManager.
     */
    public static void springbootsitesUpdateMaximumSetGen(
        com.azure.resourcemanager.springappdiscovery.SpringAppDiscoveryManager manager) {
        SpringbootsitesModel resource = manager.springbootsites().getByResourceGroupWithResponse("rgspringbootsites",
            "xrmzlavpewxtfeitghdrj", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("key9581", "fakeTokenPlaceholder")).apply();
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
