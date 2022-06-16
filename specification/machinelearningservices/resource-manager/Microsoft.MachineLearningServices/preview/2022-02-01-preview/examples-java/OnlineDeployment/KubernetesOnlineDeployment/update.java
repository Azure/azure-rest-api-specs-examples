import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.Context;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.machinelearning.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.machinelearning.models.OnlineDeploymentData;
import com.azure.resourcemanager.machinelearning.models.PartialKubernetesOnlineDeployment;
import com.azure.resourcemanager.machinelearning.models.PartialManagedServiceIdentity;
import com.azure.resourcemanager.machinelearning.models.PartialSku;
import com.azure.resourcemanager.machinelearning.models.SkuTier;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

/** Samples for OnlineDeployments Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/OnlineDeployment/KubernetesOnlineDeployment/update.json
     */
    /**
     * Sample code: Update Kubernetes Online Deployment.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void updateKubernetesOnlineDeployment(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) throws IOException {
        OnlineDeploymentData resource =
            manager
                .onlineDeployments()
                .getWithResponse("test-rg", "my-aml-workspace", "testEndpointName", "testDeploymentName", Context.NONE)
                .getValue();
        resource
            .update()
            .withTags(mapOf())
            .withIdentity(
                new PartialManagedServiceIdentity()
                    .withType(ManagedServiceIdentityType.SYSTEM_ASSIGNED)
                    .withUserAssignedIdentities(
                        mapOf(
                            "string",
                            SerializerFactory
                                .createDefaultManagementSerializerAdapter()
                                .deserialize("{}", Object.class, SerializerEncoding.JSON))))
            .withKind("string")
            .withProperties(new PartialKubernetesOnlineDeployment())
            .withSku(
                new PartialSku()
                    .withCapacity(1)
                    .withFamily("string")
                    .withName("string")
                    .withSize("string")
                    .withTier(SkuTier.FREE))
            .apply();
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
