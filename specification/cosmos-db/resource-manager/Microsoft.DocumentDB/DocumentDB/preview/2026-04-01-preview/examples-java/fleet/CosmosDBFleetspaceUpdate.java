
import com.azure.resourcemanager.cosmos.models.FleetspacePropertiesFleetspaceApiKind;
import com.azure.resourcemanager.cosmos.models.FleetspacePropertiesThroughputPoolConfiguration;
import com.azure.resourcemanager.cosmos.models.FleetspaceUpdate;

/**
 * Samples for Fleetspace Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/fleet/CosmosDBFleetspaceUpdate.json
     */
    /**
     * Sample code: CosmosDB Fleetspace Update.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBFleetspaceUpdate(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getFleetspaces().update("rg1", "fleet1", "fleetspace1",
            new FleetspaceUpdate().withFleetspaceApiKind(FleetspacePropertiesFleetspaceApiKind.NO_SQL)
                .withThroughputPoolConfiguration(new FleetspacePropertiesThroughputPoolConfiguration()
                    .withMinThroughput(100000).withMaxThroughput(1000000)),
            com.azure.core.util.Context.NONE);
    }
}
