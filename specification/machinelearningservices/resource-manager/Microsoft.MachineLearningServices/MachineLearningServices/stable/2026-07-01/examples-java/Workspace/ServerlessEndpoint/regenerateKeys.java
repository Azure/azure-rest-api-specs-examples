
import com.azure.resourcemanager.machinelearning.models.KeyType;
import com.azure.resourcemanager.machinelearning.models.RegenerateEndpointKeysRequest;

/**
 * Samples for ServerlessEndpoints RegenerateKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/ServerlessEndpoint/regenerateKeys.json
     */
    /**
     * Sample code: RegenerateKeys Workspace Serverless Endpoint.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void regenerateKeysWorkspaceServerlessEndpoint(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.serverlessEndpoints().regenerateKeys("test-rg", "my-aml-workspace", "string",
            new RegenerateEndpointKeysRequest().withKeyType(KeyType.PRIMARY).withKeyValue("fakeTokenPlaceholder"),
            com.azure.core.util.Context.NONE);
    }
}
