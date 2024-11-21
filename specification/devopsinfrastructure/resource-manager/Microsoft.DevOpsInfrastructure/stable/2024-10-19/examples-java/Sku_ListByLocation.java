
/**
 * Samples for Sku ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-10-19/Sku_ListByLocation.json
     */
    /**
     * Sample code: Sku_ListByLocation.
     * 
     * @param manager Entry point to DevOpsInfrastructureManager.
     */
    public static void
        skuListByLocation(com.azure.resourcemanager.devopsinfrastructure.DevOpsInfrastructureManager manager) {
        manager.skus().listByLocation("eastus", com.azure.core.util.Context.NONE);
    }
}
