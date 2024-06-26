import com.azure.resourcemanager.machinelearning.models.BatchEndpointDefaults;
import com.azure.resourcemanager.machinelearning.models.BatchEndpointProperties;
import com.azure.resourcemanager.machinelearning.models.EndpointAuthMode;
import com.azure.resourcemanager.machinelearning.models.ManagedServiceIdentity;
import com.azure.resourcemanager.machinelearning.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.machinelearning.models.Sku;
import com.azure.resourcemanager.machinelearning.models.SkuTier;
import com.azure.resourcemanager.machinelearning.models.UserAssignedIdentity;
import java.util.HashMap;
import java.util.Map;

/** Samples for BatchEndpoints CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/BatchEndpoint/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Batch Endpoint.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createOrUpdateBatchEndpoint(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .batchEndpoints()
            .define("testEndpointName")
            .withRegion("string")
            .withExistingWorkspace("test-rg", "my-aml-workspace")
            .withProperties(
                new BatchEndpointProperties()
                    .withAuthMode(EndpointAuthMode.AMLTOKEN)
                    .withDescription("string")
                    .withProperties(mapOf("string", "string"))
                    .withDefaults(new BatchEndpointDefaults().withDeploymentName("string")))
            .withTags(mapOf())
            .withIdentity(
                new ManagedServiceIdentity()
                    .withType(ManagedServiceIdentityType.SYSTEM_ASSIGNED)
                    .withUserAssignedIdentities(mapOf("string", new UserAssignedIdentity())))
            .withKind("string")
            .withSku(
                new Sku()
                    .withName("string")
                    .withTier(SkuTier.FREE)
                    .withSize("string")
                    .withFamily("string")
                    .withCapacity(1))
            .create();
    }

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
