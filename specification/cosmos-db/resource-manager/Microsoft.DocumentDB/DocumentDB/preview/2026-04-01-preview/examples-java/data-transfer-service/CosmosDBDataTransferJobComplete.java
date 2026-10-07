
/**
 * Samples for DataTransferJobs Complete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/data-transfer-service/CosmosDBDataTransferJobComplete.json
     */
    /**
     * Sample code: CosmosDBDataTransferJobComplete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBDataTransferJobComplete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getDataTransferJobs().completeWithResponse("rg1", "ddb1", "j1",
            com.azure.core.util.Context.NONE);
    }
}
