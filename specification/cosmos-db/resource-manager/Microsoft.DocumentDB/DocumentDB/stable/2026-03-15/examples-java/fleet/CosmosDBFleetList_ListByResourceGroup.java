
/**
 * Samples for Fleet ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/fleet/CosmosDBFleetList_ListByResourceGroup.json
     */
    /**
     * Sample code: CosmosDB Fleet List by Resource Group.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBFleetListByResourceGroup(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getFleets().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
