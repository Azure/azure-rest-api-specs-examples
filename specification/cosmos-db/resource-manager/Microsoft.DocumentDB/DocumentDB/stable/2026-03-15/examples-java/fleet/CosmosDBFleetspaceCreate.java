
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
     * x-ms-original-file: 2026-03-15/fleet/CosmosDBFleetspaceCreate.json
     */
    /**
     * Sample code: CosmosDB Fleetspace Create.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBFleetspaceCreate(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getFleetspaces().create("rg1", "fleet1", "fleetspace1",
            new FleetspaceResourceInner().withFleetspaceApiKind(FleetspacePropertiesFleetspaceApiKind.NO_SQL)
                .withServiceTier(FleetspacePropertiesServiceTier.GENERAL_PURPOSE)
                .withDataRegions(Arrays.asList("westus2"))
                .withThroughputPoolConfiguration(new FleetspacePropertiesThroughputPoolConfiguration()
                    .withMinThroughput(100000).withMaxThroughput(500000)),
            com.azure.core.util.Context.NONE);
    }
}
