
/**
 * Samples for PaymentHsmClusterPrivateLinkResources ListByPaymentHsmCluster.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-12-01-preview/PaymentHsmClusterPrivateLinkResource_ListByPaymentHsmCluster_MaximumSet_Gen.json
     */
    /**
     * Sample code: PaymentHsmClusterPrivateLinkResource_ListByPaymentHsmCluster_MaximumSet_Gen.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void paymentHsmClusterPrivateLinkResourceListByPaymentHsmClusterMaximumSetGen(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.paymentHsmClusterPrivateLinkResources().listByPaymentHsmCluster("rgpaymenthsm", "phsm1",
            com.azure.core.util.Context.NONE);
    }
}
