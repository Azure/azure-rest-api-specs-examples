
/**
 * Samples for QueueServices List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/QueueServicesList.json
     */
    /**
     * Sample code: QueueServicesList.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void queueServicesList(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getQueueServices().listWithResponse("res9290", "sto1590",
            com.azure.core.util.Context.NONE);
    }
}
