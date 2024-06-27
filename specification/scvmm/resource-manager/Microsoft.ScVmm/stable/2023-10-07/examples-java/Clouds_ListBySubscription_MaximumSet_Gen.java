
/**
 * Samples for Clouds List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * Clouds_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: Clouds_ListBySubscription_MaximumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void cloudsListBySubscriptionMaximumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.clouds().list(com.azure.core.util.Context.NONE);
    }
}
