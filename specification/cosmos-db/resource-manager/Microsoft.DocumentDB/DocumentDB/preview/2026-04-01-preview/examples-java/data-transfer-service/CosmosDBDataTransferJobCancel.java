
/**
 * Samples for DataTransferJobs Cancel.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/data-transfer-service/CosmosDBDataTransferJobCancel.json
     */
    /**
     * Sample code: CosmosDBDataTransferJobCancel.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBDataTransferJobCancel(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getDataTransferJobs().cancelWithResponse("rg1", "ddb1", "j1",
            com.azure.core.util.Context.NONE);
    }
}
