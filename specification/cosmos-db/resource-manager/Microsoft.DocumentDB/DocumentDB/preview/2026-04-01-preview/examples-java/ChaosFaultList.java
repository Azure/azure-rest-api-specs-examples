
/**
 * Samples for ChaosFault List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/ChaosFaultList.json
     */
    /**
     * Sample code: ChaosFaultList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void chaosFaultList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getChaosFaults().list("rg1", "ddb1", com.azure.core.util.Context.NONE);
    }
}
