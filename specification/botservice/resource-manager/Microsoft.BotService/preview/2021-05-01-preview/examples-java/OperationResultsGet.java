
/**
 * Samples for OperationResults Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/
     * OperationResultsGet.json
     */
    /**
     * Sample code: Get operation result.
     * 
     * @param manager Entry point to BotServiceManager.
     */
    public static void getOperationResult(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.operationResults().get("exampleid", com.azure.core.util.Context.NONE);
    }
}
