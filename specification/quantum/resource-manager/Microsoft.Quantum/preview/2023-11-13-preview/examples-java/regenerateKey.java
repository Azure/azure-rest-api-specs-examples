
import com.azure.resourcemanager.quantum.models.ApiKeys;
import com.azure.resourcemanager.quantum.models.KeyType;
import java.util.Arrays;

/**
 * Samples for WorkspaceOperation RegenerateKeys.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/quantum/resource-manager/Microsoft.Quantum/preview/2023-11-13-preview/examples/regenerateKey.json
     */
    /**
     * Sample code: RegenerateKey.
     * 
     * @param manager Entry point to AzureQuantumManager.
     */
    public static void regenerateKey(com.azure.resourcemanager.quantum.AzureQuantumManager manager) {
        manager.workspaceOperations().regenerateKeysWithResponse("quantumResourcegroup", "quantumworkspace1",
            new ApiKeys().withKeys(Arrays.asList(KeyType.PRIMARY, KeyType.SECONDARY)),
            com.azure.core.util.Context.NONE);
    }
}
