
import com.azure.resourcemanager.fluidrelay.models.Identity;
import com.azure.resourcemanager.fluidrelay.models.ResourceIdentityType;
import com.azure.resourcemanager.fluidrelay.models.StorageSku;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for FluidRelayServers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/fluidrelay/resource-manager/Microsoft.FluidRelay/stable/2022-06-01/examples/
     * FluidRelayServers_CreateOrUpdate.json
     */
    /**
     * Sample code: Create a Fluid Relay server.
     * 
     * @param manager Entry point to FluidRelayManager.
     */
    public static void createAFluidRelayServer(com.azure.resourcemanager.fluidrelay.FluidRelayManager manager) {
        manager.fluidRelayServers().define("myFluidRelayServer").withRegion("west-us")
            .withExistingResourceGroup("myResourceGroup").withTags(mapOf("Category", "sales"))
            .withIdentity(new Identity().withType(ResourceIdentityType.SYSTEM_ASSIGNED))
            .withStoragesku(StorageSku.BASIC).create();
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
