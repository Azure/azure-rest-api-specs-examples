
/**
 * Samples for ChaosFault Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/ChaosFaultGet.json
     */
    /**
     * Sample code: ChaosFaultGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void chaosFaultGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getChaosFaults().getWithResponse("rg1", "ddb1", "ServiceUnavailability",
            com.azure.core.util.Context.NONE);
    }
}
