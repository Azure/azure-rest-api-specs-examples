
/**
 * Samples for Service Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBGraphAPIComputeServiceDelete.json
     */
    /**
     * Sample code: GraphAPIComputeServiceDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void graphAPIComputeServiceDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getServices().delete("rg1", "ddb1", "GraphAPICompute",
            com.azure.core.util.Context.NONE);
    }
}
