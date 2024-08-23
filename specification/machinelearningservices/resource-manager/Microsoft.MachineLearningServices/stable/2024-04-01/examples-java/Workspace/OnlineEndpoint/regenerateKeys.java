
import com.azure.resourcemanager.machinelearning.models.KeyType;
import com.azure.resourcemanager.machinelearning.models.RegenerateEndpointKeysRequest;

/**
 * Samples for OnlineEndpoints RegenerateKeys.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/OnlineEndpoint/regenerateKeys.json
     */
    /**
     * Sample code: RegenerateKeys Workspace Online Endpoint.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void regenerateKeysWorkspaceOnlineEndpoint(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.onlineEndpoints().regenerateKeys("test-rg", "my-aml-workspace", "testEndpointName",
            new RegenerateEndpointKeysRequest().withKeyType(KeyType.PRIMARY).withKeyValue("fakeTokenPlaceholder"),
            com.azure.core.util.Context.NONE);
    }
}
