
/**
 * Samples for CollectorPolicies List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2022-11-01/CollectorPoliciesList.json
     */
    /**
     * Sample code: List of Collection Policies.
     * 
     * @param manager Entry point to AzureTrafficCollectorManager.
     */
    public static void
        listOfCollectionPolicies(com.azure.resourcemanager.networkfunction.AzureTrafficCollectorManager manager) {
        manager.collectorPolicies().list("rg1", "atc", com.azure.core.util.Context.NONE);
    }
}
