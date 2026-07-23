
import com.azure.resourcemanager.cosmos.models.FleetspacePropertiesFleetspaceApiKind;
import com.azure.resourcemanager.cosmos.models.FleetspacePropertiesThroughputPoolConfiguration;
import com.azure.resourcemanager.cosmos.models.FleetspaceUpdate;
import java.util.Arrays;

/**
 * Samples for Fleetspace Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/fleet/CosmosDBFleetspaceUpdate.json
     */
    /**
     * Sample code: CosmosDB Fleetspace Update.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBFleetspaceUpdate(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getFleetspaces().update("rg1", "fleet1", "fleetspace1", new FleetspaceUpdate()
            .withFleetspaceApiKind(FleetspacePropertiesFleetspaceApiKind.NO_SQL)
            .withDataRegions(Arrays.asList("westus2")).withThroughputPoolConfiguration(
                new FleetspacePropertiesThroughputPoolConfiguration().withMinThroughput(3000).withMaxThroughput(4000)),
            com.azure.core.util.Context.NONE);
    }
}
