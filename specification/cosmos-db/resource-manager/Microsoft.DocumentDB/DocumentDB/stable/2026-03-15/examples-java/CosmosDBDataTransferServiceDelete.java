
/**
 * Samples for Service Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBDataTransferServiceDelete.json
     */
    /**
     * Sample code: DataTransferServiceDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void dataTransferServiceDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getServices().delete("rg1", "ddb1", "DataTransfer", com.azure.core.util.Context.NONE);
    }
}
