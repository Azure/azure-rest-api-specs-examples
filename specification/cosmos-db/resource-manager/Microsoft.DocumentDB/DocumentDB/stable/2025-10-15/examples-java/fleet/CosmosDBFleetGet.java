
/**
 * Samples for Fleet GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/fleet/
     * CosmosDBFleetGet.json
     */
    /**
     * Sample code: CosmosDB Fleet Get.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBFleetGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getFleets().getByResourceGroupWithResponse("rg1", "fleet1",
            com.azure.core.util.Context.NONE);
    }
}
