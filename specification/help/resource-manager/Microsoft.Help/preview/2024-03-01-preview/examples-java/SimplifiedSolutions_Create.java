
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for SimplifiedSolutions Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/help/resource-manager/Microsoft.Help/preview/2024-03-01-preview/examples/SimplifiedSolutions_Create
     * .json
     */
    /**
     * Sample code: Solution_Create.
     * 
     * @param manager Entry point to SelfHelpManager.
     */
    public static void solutionCreate(com.azure.resourcemanager.selfhelp.SelfHelpManager manager) {
        manager.simplifiedSolutions().define("simplifiedSolutionsResourceName1").withExistingScope(
            "subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp")
            .withSolutionId("sampleSolutionId")
            .withParameters(mapOf("resourceUri",
                "subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp"))
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
