
import com.azure.resourcemanager.machinelearning.models.ContentSafety;
import com.azure.resourcemanager.machinelearning.models.ContentSafetyStatus;
import com.azure.resourcemanager.machinelearning.models.ManagedServiceIdentity;
import com.azure.resourcemanager.machinelearning.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.machinelearning.models.ModelSettings;
import com.azure.resourcemanager.machinelearning.models.ServerlessEndpointProperties;
import com.azure.resourcemanager.machinelearning.models.ServerlessInferenceEndpointAuthMode;
import com.azure.resourcemanager.machinelearning.models.Sku;
import com.azure.resourcemanager.machinelearning.models.SkuTier;
import com.azure.resourcemanager.machinelearning.models.UserAssignedIdentity;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ServerlessEndpoints CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/ServerlessEndpoint/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Workspace Serverless Endpoint.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createOrUpdateWorkspaceServerlessEndpoint(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.serverlessEndpoints().define("string").withRegion("string")
            .withExistingWorkspace("test-rg", "my-aml-workspace")
            .withProperties(
                new ServerlessEndpointProperties().withModelSettings(new ModelSettings().withModelId("string"))
                    .withAuthMode(ServerlessInferenceEndpointAuthMode.KEY)
                    .withContentSafety(new ContentSafety().withContentSafetyStatus(ContentSafetyStatus.ENABLED)))
            .withTags(mapOf()).withKind("string")
            .withIdentity(new ManagedServiceIdentity().withType(ManagedServiceIdentityType.SYSTEM_ASSIGNED)
                .withUserAssignedIdentities(mapOf("string", new UserAssignedIdentity())))
            .withSku(new Sku().withName("string").withTier(SkuTier.STANDARD).withSize("string").withFamily("string")
                .withCapacity(1))
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
