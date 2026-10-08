
/**
 * Samples for DataTransferJobs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/data-transfer-service/CosmosDBDataTransferJobGet.json
     */
    /**
     * Sample code: CosmosDBDataTransferJobGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBDataTransferJobGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getDataTransferJobs().getWithResponse("rg1", "ddb1", "j1",
            com.azure.core.util.Context.NONE);
    }
}
