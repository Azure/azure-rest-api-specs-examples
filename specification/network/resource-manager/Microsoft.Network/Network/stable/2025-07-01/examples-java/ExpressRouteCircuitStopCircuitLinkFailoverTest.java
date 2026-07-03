
import com.azure.resourcemanager.network.models.ExpressRouteLinkFailoverStopApiParameters;

/**
 * Samples for ExpressRouteCircuits StopCircuitLinkFailoverTest.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCircuitStopCircuitLinkFailoverTest.json
     */
    /**
     * Sample code: ExpressRouteCircuitStopCircuitLinkFailoverTest.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        expressRouteCircuitStopCircuitLinkFailoverTest(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCircuits().stopCircuitLinkFailoverTest("rg1", "circuit1",
            new ExpressRouteLinkFailoverStopApiParameters().withCircuitTestCategory("BgpDisconnect")
                .withLinkType("Primary").withWasSimulationSuccessful(true).withIsVerified(true),
            com.azure.core.util.Context.NONE);
    }
}
