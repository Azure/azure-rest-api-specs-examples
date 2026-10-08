
/**
 * Samples for DataTransferJobs Resume.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/data-transfer-service/CosmosDBDataTransferJobResume.json
     */
    /**
     * Sample code: CosmosDBDataTransferJobResume.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBDataTransferJobResume(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getDataTransferJobs().resumeWithResponse("rg1", "ddb1", "j1",
            com.azure.core.util.Context.NONE);
    }
}
