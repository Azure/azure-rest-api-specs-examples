
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Springbootsites CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/
     * examples/springbootsites_CreateOrUpdate_MinimumSet_Gen.json
     */
    /**
     * Sample code: springbootsites_CreateOrUpdate_MinimumSet_Gen.
     * 
     * @param manager Entry point to SpringAppDiscoveryManager.
     */
    public static void springbootsitesCreateOrUpdateMinimumSetGen(
        com.azure.resourcemanager.springappdiscovery.SpringAppDiscoveryManager manager) {
        manager.springbootsites().define("xrmzlavpewxtfeitghdrj").withRegion("tgobtvxktootwhhvjtsmpddvlqlrq")
            .withExistingResourceGroup("rgspringbootsites").create();
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
