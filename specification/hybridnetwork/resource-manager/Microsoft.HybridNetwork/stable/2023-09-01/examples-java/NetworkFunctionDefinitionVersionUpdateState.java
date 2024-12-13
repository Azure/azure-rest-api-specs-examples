
import com.azure.resourcemanager.hybridnetwork.fluent.models.NetworkFunctionDefinitionVersionUpdateStateInner;
import com.azure.resourcemanager.hybridnetwork.models.VersionState;

/**
 * Samples for NetworkFunctionDefinitionVersions UpdateState.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * NetworkFunctionDefinitionVersionUpdateState.json
     */
    /**
     * Sample code: Update network function definition version state.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void updateNetworkFunctionDefinitionVersionState(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.networkFunctionDefinitionVersions().updateState("rg", "TestPublisher", "TestSkuGroup", "1.0.0",
            new NetworkFunctionDefinitionVersionUpdateStateInner().withVersionState(VersionState.ACTIVE),
            com.azure.core.util.Context.NONE);
    }
}
