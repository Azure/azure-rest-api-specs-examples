import com.azure.core.util.Context;

/** Samples for RestorableDatabaseAccounts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBRestorableDatabaseAccountNoLocationList.json
     */
    /**
     * Sample code: CosmosDBRestorableDatabaseAccountNoLocationList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBRestorableDatabaseAccountNoLocationList(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getRestorableDatabaseAccounts().list(Context.NONE);
    }
}
