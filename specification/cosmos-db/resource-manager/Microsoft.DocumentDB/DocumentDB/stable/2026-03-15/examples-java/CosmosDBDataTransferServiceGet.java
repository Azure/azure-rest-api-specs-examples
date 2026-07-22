
/**
 * Samples for Service Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBDataTransferServiceGet.json
     */
    /**
     * Sample code: DataTransferServiceGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void dataTransferServiceGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getServices().getWithResponse("rg1", "ddb1", "DataTransfer",
            com.azure.core.util.Context.NONE);
    }
}
