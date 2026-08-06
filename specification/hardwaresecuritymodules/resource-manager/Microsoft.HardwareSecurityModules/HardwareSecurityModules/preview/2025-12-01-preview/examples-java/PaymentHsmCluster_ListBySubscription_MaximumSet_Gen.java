
/**
 * Samples for PaymentHsmClusters List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01-preview/PaymentHsmCluster_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: PaymentHsmCluster_ListBySubscription_MaximumSet_Gen.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void paymentHsmClusterListBySubscriptionMaximumSetGen(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.paymentHsmClusters().list(null, com.azure.core.util.Context.NONE);
    }
}
