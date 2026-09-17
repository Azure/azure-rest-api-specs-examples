
import com.azure.resourcemanager.network.models.MigrateExpressRouteCircuitRequest;
import com.azure.resourcemanager.network.models.PortMapping;
import java.util.Arrays;

/**
 * Samples for ExpressRouteCrossConnections RollbackCircuitMigration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/ExpressRouteCrossConnectionRollbackCircuitMigration.json
     */
    /**
     * Sample code: RollbackExpressRouteCircuitMigration.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void rollbackExpressRouteCircuitMigration(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCrossConnections().rollbackCircuitMigration(
            "CrossConnection-SiliconValley", "<circuitServiceKey>",
            new MigrateExpressRouteCircuitRequest().withTargetPeeringLocation("SiliconValley")
                .withTargetPortMapping(
                    Arrays.asList(new PortMapping().withSourcePortId("sourcePort1").withTargetPortId("targetPort1"),
                        new PortMapping().withSourcePortId("sourcePort2").withTargetPortId("targetPort2")))
                .withPortId("sourcePort1"),
            com.azure.core.util.Context.NONE);
    }
}
