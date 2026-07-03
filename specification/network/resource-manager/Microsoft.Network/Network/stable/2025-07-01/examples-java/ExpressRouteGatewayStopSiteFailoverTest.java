
import com.azure.resourcemanager.network.models.ExpressRouteFailoverStopApiParameters;
import com.azure.resourcemanager.network.models.FailoverConnectionDetails;
import java.util.Arrays;

/**
 * Samples for ExpressRouteGateways StopSiteFailoverTest.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteGatewayStopSiteFailoverTest.json
     */
    /**
     * Sample code: ExpressRouteGatewayStopSiteFailoverTest.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        expressRouteGatewayStopSiteFailoverTest(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteGateways().stopSiteFailoverTest("rg1", "ergw1",
            new ExpressRouteFailoverStopApiParameters().withPeeringLocation("Vancouver")
                .withWasSimulationSuccessful(true).withDetails(Arrays.asList(new FailoverConnectionDetails()
                    .withFailoverConnectionName("conn1").withFailoverLocation("Denver").withIsVerified(true))),
            com.azure.core.util.Context.NONE);
    }
}
