
/**
 * Samples for DataTransferJobs ListByDatabaseAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/data-transfer-service/CosmosDBDataTransferJobFeed.json
     */
    /**
     * Sample code: CosmosDBDataTransferJobFeed.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBDataTransferJobFeed(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getDataTransferJobs().listByDatabaseAccount("rg1", "ddb1",
            com.azure.core.util.Context.NONE);
    }
}
