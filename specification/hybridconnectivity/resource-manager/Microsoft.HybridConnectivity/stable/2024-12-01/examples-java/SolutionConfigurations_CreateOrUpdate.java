
import com.azure.resourcemanager.hybridconnectivity.models.SolutionConfigurationProperties;
import com.azure.resourcemanager.hybridconnectivity.models.SolutionSettings;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for SolutionConfigurations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-12-01/SolutionConfigurations_CreateOrUpdate.json
     */
    /**
     * Sample code: SolutionConfigurations_CreateOrUpdate.
     * 
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void solutionConfigurationsCreateOrUpdate(
        com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        manager.solutionConfigurations().define("keebwujt").withExistingResourceUri("ymuj")
            .withProperties(new SolutionConfigurationProperties().withSolutionType("nmtqllkyohwtsthxaimsye")
                .withSolutionSettings(new SolutionSettings().withAdditionalProperties(mapOf())))
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
