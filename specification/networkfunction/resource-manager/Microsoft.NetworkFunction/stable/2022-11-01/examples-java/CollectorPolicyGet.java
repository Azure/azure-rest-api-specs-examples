import com.azure.core.util.Context;

/** Samples for CollectorPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkfunction/resource-manager/Microsoft.NetworkFunction/stable/2022-11-01/examples/CollectorPolicyGet.json
     */
    /**
     * Sample code: Get Collection Policy.
     *
     * @param manager Entry point to AzureTrafficCollectorManager.
     */
    public static void getCollectionPolicy(
        com.azure.resourcemanager.networkfunction.AzureTrafficCollectorManager manager) {
        manager.collectorPolicies().getWithResponse("rg1", "atc", "cp1", Context.NONE);
    }
}
