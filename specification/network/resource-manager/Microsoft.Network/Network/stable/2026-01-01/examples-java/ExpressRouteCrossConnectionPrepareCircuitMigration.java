
import com.azure.resourcemanager.network.models.MigrateExpressRouteCircuitRequest;
import com.azure.resourcemanager.network.models.PortMapping;
import java.util.Arrays;

/**
 * Samples for ExpressRouteCrossConnections PrepareCircuitMigration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/ExpressRouteCrossConnectionPrepareCircuitMigration.json
     */
    /**
     * Sample code: PrepareExpressRouteCircuitMigration.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void prepareExpressRouteCircuitMigration(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCrossConnections().prepareCircuitMigration(
            "CrossConnection-SiliconValley", "<circuitServiceKey>",
            new MigrateExpressRouteCircuitRequest().withTargetPeeringLocation("SiliconValley")
                .withTargetPortMapping(
                    Arrays.asList(new PortMapping().withSourcePortId("sourcePort1").withTargetPortId("targetPort1"),
                        new PortMapping().withSourcePortId("sourcePort2").withTargetPortId("targetPort2")))
                .withPortId("sourcePort1"),
            com.azure.core.util.Context.NONE);
    }
}
