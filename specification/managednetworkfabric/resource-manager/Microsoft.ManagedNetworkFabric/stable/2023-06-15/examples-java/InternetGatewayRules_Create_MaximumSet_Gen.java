
import com.azure.resourcemanager.managednetworkfabric.models.Action;
import com.azure.resourcemanager.managednetworkfabric.models.RuleProperties;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for InternetGatewayRules Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/
     * InternetGatewayRules_Create_MaximumSet_Gen.json
     */
    /**
     * Sample code: InternetGatewayRules_Create_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void internetGatewayRulesCreateMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.internetGatewayRules().define("example-internetGatewayRule").withRegion("eastus")
            .withExistingResourceGroup("example-rg")
            .withRuleProperties(
                new RuleProperties().withAction(Action.ALLOW).withAddressList(Arrays.asList("10.10.10.10")))
            .withTags(mapOf("keyID", "fakeTokenPlaceholder")).withAnnotation("annotationValue").create();
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
