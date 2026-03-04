
/**
 * Samples for CollectorPolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2022-11-01/CollectorPolicyDelete.json
     */
    /**
     * Sample code: Delete Collection Policy.
     * 
     * @param manager Entry point to AzureTrafficCollectorManager.
     */
    public static void
        deleteCollectionPolicy(com.azure.resourcemanager.networkfunction.AzureTrafficCollectorManager manager) {
        manager.collectorPolicies().delete("rg1", "atc", "cp1", com.azure.core.util.Context.NONE);
    }
}
