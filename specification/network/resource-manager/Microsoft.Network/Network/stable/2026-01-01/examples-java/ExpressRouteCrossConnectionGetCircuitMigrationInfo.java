
import com.azure.resourcemanager.network.models.MigrateExpressRouteCircuitValidateAndHealthCheckRequest;
import com.azure.resourcemanager.network.models.PortMapping;
import java.util.Arrays;

/**
 * Samples for ExpressRouteCrossConnections GetCircuitMigrationInfo.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/ExpressRouteCrossConnectionGetCircuitMigrationInfo.json
     */
    /**
     * Sample code: GetExpressRouteCircuitMigrationInfo.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getExpressRouteCircuitMigrationInfo(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCrossConnections().getCircuitMigrationInfo(
            "CrossConnection-SiliconValley", "<circuitServiceKey>",
            new MigrateExpressRouteCircuitValidateAndHealthCheckRequest().withTargetPeeringLocation("SiliconValley")
                .withTargetPortMapping(
                    Arrays.asList(new PortMapping().withSourcePortId("sourcePort1").withTargetPortId("targetPort1"),
                        new PortMapping().withSourcePortId("sourcePort2").withTargetPortId("targetPort2"))),
            com.azure.core.util.Context.NONE);
    }
}
