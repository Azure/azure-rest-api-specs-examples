import com.azure.core.util.Context;

/** Samples for CollectorPolicies Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkfunction/resource-manager/Microsoft.NetworkFunction/stable/2022-05-01/examples/CollectorPolicyDelete.json
     */
    /**
     * Sample code: Delete Collection Policy.
     *
     * @param manager Entry point to AzureTrafficCollectorManager.
     */
    public static void deleteCollectionPolicy(
        com.azure.resourcemanager.networkfunction.AzureTrafficCollectorManager manager) {
        manager.collectorPolicies().delete("rg1", "atc", "cp1", Context.NONE);
    }
}
