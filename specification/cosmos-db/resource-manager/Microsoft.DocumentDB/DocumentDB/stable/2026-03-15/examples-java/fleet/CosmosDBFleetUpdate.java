
import com.azure.resourcemanager.cosmos.models.FleetResourceUpdate;

/**
 * Samples for Fleet Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/fleet/CosmosDBFleetUpdate.json
     */
    /**
     * Sample code: CosmosDB Fleet Update.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBFleetUpdate(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getFleets().updateWithResponse("rg1", "fleet1", new FleetResourceUpdate(),
            com.azure.core.util.Context.NONE);
    }
}
