
import com.azure.resourcemanager.cosmos.models.FleetspacePropertiesFleetspaceApiKind;
import com.azure.resourcemanager.cosmos.models.FleetspacePropertiesThroughputPoolConfiguration;
import com.azure.resourcemanager.cosmos.models.FleetspaceUpdate;
import java.util.Arrays;

/**
 * Samples for Fleetspace Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/fleet/
     * CosmosDBFleetspaceUpdate.json
     */
    /**
     * Sample code: CosmosDB Fleetspace Update.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBFleetspaceUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getFleetspaces().update("rg1", "fleet1", "fleetspace1",
            new FleetspaceUpdate().withFleetspaceApiKind(FleetspacePropertiesFleetspaceApiKind.NO_SQL)
                .withDataRegions(Arrays.asList("westus2"))
                .withThroughputPoolConfiguration(new FleetspacePropertiesThroughputPoolConfiguration()
                    .withMinThroughput(3000).withMaxThroughput(4000)),
            com.azure.core.util.Context.NONE);
    }
}
