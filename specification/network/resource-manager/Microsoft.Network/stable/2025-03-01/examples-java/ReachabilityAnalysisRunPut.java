
import com.azure.resourcemanager.network.fluent.models.ReachabilityAnalysisRunInner;
import com.azure.resourcemanager.network.models.ReachabilityAnalysisRunProperties;

/**
 * Samples for ReachabilityAnalysisRuns Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/ReachabilityAnalysisRunPut.
     * json
     */
    /**
     * Sample code: ReachabilityAnalysisRunCreate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void reachabilityAnalysisRunCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getReachabilityAnalysisRuns().createWithResponse("rg1",
            "testNetworkManager", "testWorkspace", "testAnalysisRunName",
            new ReachabilityAnalysisRunInner().withProperties(new ReachabilityAnalysisRunProperties()
                .withDescription("A sample reachability analysis run").withIntentId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager/verifierWorkspaces/testVerifierWorkspace1/reachabilityAnalysisIntents/testReachabilityAnalysisIntenant1")),
            com.azure.core.util.Context.NONE);
    }
}
