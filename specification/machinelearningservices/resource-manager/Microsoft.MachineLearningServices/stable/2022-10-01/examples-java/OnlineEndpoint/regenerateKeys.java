import com.azure.core.util.Context;
import com.azure.resourcemanager.machinelearning.models.KeyType;
import com.azure.resourcemanager.machinelearning.models.RegenerateEndpointKeysRequest;

/** Samples for OnlineEndpoints RegenerateKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/OnlineEndpoint/regenerateKeys.json
     */
    /**
     * Sample code: RegenerateKeys Online Endpoint.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void regenerateKeysOnlineEndpoint(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .onlineEndpoints()
            .regenerateKeys(
                "test-rg",
                "my-aml-workspace",
                "testEndpointName",
                new RegenerateEndpointKeysRequest().withKeyType(KeyType.PRIMARY).withKeyValue("fakeTokenPlaceholder"),
                Context.NONE);
    }
}
