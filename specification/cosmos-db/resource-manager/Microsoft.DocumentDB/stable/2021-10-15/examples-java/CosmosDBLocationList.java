import com.azure.core.util.Context;

/** Samples for Locations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBLocationList.json
     */
    /**
     * Sample code: CosmosDBLocationList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBLocationList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getLocations().list(Context.NONE);
    }
}
