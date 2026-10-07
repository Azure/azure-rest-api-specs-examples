
/**
 * Samples for DataTransferJobs Pause.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/data-transfer-service/CosmosDBDataTransferJobPause.json
     */
    /**
     * Sample code: CosmosDBDataTransferJobPause.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBDataTransferJobPause(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getDataTransferJobs().pauseWithResponse("rg1", "ddb1", "j1",
            com.azure.core.util.Context.NONE);
    }
}
