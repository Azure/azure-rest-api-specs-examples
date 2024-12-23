
import com.azure.resourcemanager.synapse.models.IntegrationRuntimeAuthKeyName;
import com.azure.resourcemanager.synapse.models.IntegrationRuntimeRegenerateKeyParameters;

/**
 * Samples for IntegrationRuntimeAuthKeysOperation Regenerate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-03-01/examples/
     * IntegrationRuntimes_RegenerateAuthKey.json
     */
    /**
     * Sample code: Regenerate auth key.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void regenerateAuthKey(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.integrationRuntimeAuthKeysOperations().regenerateWithResponse("exampleResourceGroup",
            "exampleWorkspace", "exampleIntegrationRuntime",
            new IntegrationRuntimeRegenerateKeyParameters().withKeyName(IntegrationRuntimeAuthKeyName.AUTH_KEY2),
            com.azure.core.util.Context.NONE);
    }
}
