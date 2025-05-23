
import com.azure.resourcemanager.containerservicefleet.models.AgentProfile;
import com.azure.resourcemanager.containerservicefleet.models.FleetHubProfile;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Fleets Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/Fleets_CreateOrUpdate.json
     */
    /**
     * Sample code: Creates a Fleet resource with a long running operation.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void createsAFleetResourceWithALongRunningOperation(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.fleets().define("fleet1").withRegion("East US").withExistingResourceGroup("rg1")
            .withTags(mapOf("tier", "production", "archv2", "")).withHubProfile(new FleetHubProfile()
                .withDnsPrefix("dnsprefix1").withAgentProfile(new AgentProfile().withVmSize("Standard_DS1")))
            .create();
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
