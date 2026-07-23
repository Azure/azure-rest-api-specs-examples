
/**
 * Samples for Service Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBGraphAPIComputeServiceGet.json
     */
    /**
     * Sample code: GraphAPIComputeServiceGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void graphAPIComputeServiceGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getServices().getWithResponse("rg1", "ddb1", "GraphAPICompute",
            com.azure.core.util.Context.NONE);
    }
}
