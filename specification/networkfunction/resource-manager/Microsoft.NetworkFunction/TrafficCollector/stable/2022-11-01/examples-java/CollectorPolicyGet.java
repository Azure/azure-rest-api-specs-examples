
/**
 * Samples for CollectorPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2022-11-01/CollectorPolicyGet.json
     */
    /**
     * Sample code: Get Collection Policy.
     * 
     * @param manager Entry point to AzureTrafficCollectorManager.
     */
    public static void
        getCollectionPolicy(com.azure.resourcemanager.networkfunction.AzureTrafficCollectorManager manager) {
        manager.collectorPolicies().getWithResponse("rg1", "atc", "cp1", com.azure.core.util.Context.NONE);
    }
}
