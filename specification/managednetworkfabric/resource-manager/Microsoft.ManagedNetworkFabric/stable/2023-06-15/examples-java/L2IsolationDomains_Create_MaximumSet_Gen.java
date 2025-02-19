
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for L2IsolationDomains Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/
     * L2IsolationDomains_Create_MaximumSet_Gen.json
     */
    /**
     * Sample code: L2IsolationDomains_Create_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void l2IsolationDomainsCreateMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.l2IsolationDomains().define("example-l2domain").withRegion("eastus")
            .withExistingResourceGroup("example-rg")
            .withNetworkFabricId(
                "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric")
            .withVlanId(501).withTags(mapOf("keyID", "fakeTokenPlaceholder")).withMtu(1500).withAnnotation("annotation")
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
