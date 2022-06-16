import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.models.IntegrationRuntimeAuthKeyName;
import com.azure.resourcemanager.synapse.models.IntegrationRuntimeRegenerateKeyParameters;

/** Samples for IntegrationRuntimeAuthKeysOperation Regenerate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_RegenerateAuthKey.json
     */
    /**
     * Sample code: Regenerate auth key.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void regenerateAuthKey(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .integrationRuntimeAuthKeysOperations()
            .regenerateWithResponse(
                "exampleResourceGroup",
                "exampleWorkspace",
                "exampleIntegrationRuntime",
                new IntegrationRuntimeRegenerateKeyParameters().withKeyName(IntegrationRuntimeAuthKeyName.AUTH_KEY2),
                Context.NONE);
    }
}
