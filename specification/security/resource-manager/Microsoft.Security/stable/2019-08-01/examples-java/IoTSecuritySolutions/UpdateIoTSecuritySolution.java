
import com.azure.resourcemanager.security.models.IoTSecuritySolutionModel;
import com.azure.resourcemanager.security.models.RecommendationConfigStatus;
import com.azure.resourcemanager.security.models.RecommendationConfigurationProperties;
import com.azure.resourcemanager.security.models.RecommendationType;
import com.azure.resourcemanager.security.models.UserDefinedResourcesProperties;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for IotSecuritySolution Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/
     * UpdateIoTSecuritySolution.json
     */
    /**
     * Sample code: Use this method to update existing IoT Security solution.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void
        useThisMethodToUpdateExistingIoTSecuritySolution(com.azure.resourcemanager.security.SecurityManager manager) {
        IoTSecuritySolutionModel resource = manager.iotSecuritySolutions()
            .getByResourceGroupWithResponse("myRg", "default", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("foo", "bar"))
            .withUserDefinedResources(new UserDefinedResourcesProperties()
                .withQuery("where type != \"microsoft.devices/iothubs\" | where name contains \"v2\"")
                .withQuerySubscriptions(Arrays.asList("075423e9-7d33-4166-8bdf-3920b04e3735")))
            .withRecommendationsConfiguration(Arrays.asList(
                new RecommendationConfigurationProperties().withRecommendationType(RecommendationType.IO_T_OPEN_PORTS)
                    .withStatus(RecommendationConfigStatus.DISABLED),
                new RecommendationConfigurationProperties()
                    .withRecommendationType(RecommendationType.IO_T_SHARED_CREDENTIALS)
                    .withStatus(RecommendationConfigStatus.DISABLED)))
            .apply();
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
