import com.azure.resourcemanager.machinelearning.models.EndpointAuthMode;
import com.azure.resourcemanager.machinelearning.models.ManagedServiceIdentity;
import com.azure.resourcemanager.machinelearning.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.machinelearning.models.OnlineEndpointDetails;
import com.azure.resourcemanager.machinelearning.models.Sku;
import com.azure.resourcemanager.machinelearning.models.SkuTier;
import com.azure.resourcemanager.machinelearning.models.UserAssignedIdentity;
import java.util.HashMap;
import java.util.Map;

/** Samples for OnlineEndpoints CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/OnlineEndpoint/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Online Endpoint.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void createOrUpdateOnlineEndpoint(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .onlineEndpoints()
            .define("testEndpointName")
            .withRegion("string")
            .withExistingWorkspace("test-rg", "my-aml-workspace")
            .withProperties(
                new OnlineEndpointDetails()
                    .withAuthMode(EndpointAuthMode.AMLTOKEN)
                    .withDescription("string")
                    .withProperties(mapOf("string", "string"))
                    .withCompute("string")
                    .withTraffic(mapOf("string", 1)))
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
