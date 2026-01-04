
import com.azure.resourcemanager.cosmos.fluent.models.FleetspaceResourceInner;
import com.azure.resourcemanager.cosmos.models.FleetspacePropertiesFleetspaceApiKind;
import com.azure.resourcemanager.cosmos.models.FleetspacePropertiesServiceTier;
import com.azure.resourcemanager.cosmos.models.FleetspacePropertiesThroughputPoolConfiguration;
import java.util.Arrays;

/**
 * Samples for Fleetspace Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/fleet/
     * CosmosDBFleetspaceCreate.json
     */
    /**
     * Sample code: CosmosDB Fleetspace Create.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBFleetspaceCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getFleetspaces().create("rg1", "fleet1", "fleetspace1",
            new FleetspaceResourceInner().withFleetspaceApiKind(FleetspacePropertiesFleetspaceApiKind.NO_SQL)
                .withServiceTier(FleetspacePropertiesServiceTier.GENERAL_PURPOSE)
                .withDataRegions(Arrays.asList("westus2"))
                .withThroughputPoolConfiguration(new FleetspacePropertiesThroughputPoolConfiguration()
                    .withMinThroughput(100000).withMaxThroughput(500000)),
            com.azure.core.util.Context.NONE);
    }
}
