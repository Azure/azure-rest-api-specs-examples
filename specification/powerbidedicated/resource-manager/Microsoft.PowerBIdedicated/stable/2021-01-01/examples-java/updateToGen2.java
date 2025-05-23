
import com.azure.resourcemanager.powerbidedicated.models.CapacitySku;
import com.azure.resourcemanager.powerbidedicated.models.CapacitySkuTier;
import com.azure.resourcemanager.powerbidedicated.models.DedicatedCapacity;
import com.azure.resourcemanager.powerbidedicated.models.Mode;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Capacities Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/
     * updateToGen2.json
     */
    /**
     * Sample code: Update capacity to Generation 2.
     * 
     * @param manager Entry point to PowerBIDedicatedManager.
     */
    public static void
        updateCapacityToGeneration2(com.azure.resourcemanager.powerbidedicated.PowerBIDedicatedManager manager) {
        DedicatedCapacity resource = manager.capacities()
            .getByResourceGroupWithResponse("TestRG", "azsdktest", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("testKey", "fakeTokenPlaceholder"))
            .withSku(new CapacitySku().withName("A1").withTier(CapacitySkuTier.PBIE_AZURE)).withMode(Mode.GEN2).apply();
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
