
import com.azure.resourcemanager.hybridconnectivity.models.GenerateAwsTemplateRequest;
import com.azure.resourcemanager.hybridconnectivity.models.SolutionSettings;
import com.azure.resourcemanager.hybridconnectivity.models.SolutionTypeSettings;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for GenerateAwsTemplate Post.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-12-01/GenerateAwsTemplate_Post.json
     */
    /**
     * Sample code: GenerateAwsTemplate_Post.
     * 
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void
        generateAwsTemplatePost(com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        manager.generateAwsTemplates().postWithResponse(
            new GenerateAwsTemplateRequest().withConnectorId("pnxcfjidglabnwxit")
                .withSolutionTypes(Arrays.asList(new SolutionTypeSettings().withSolutionType("hjyownzpfxwiufmd")
                    .withSolutionSettings(new SolutionSettings().withAdditionalProperties(mapOf())))),
            com.azure.core.util.Context.NONE);
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
