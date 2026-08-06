
/**
 * Samples for PaymentHsmClusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01-preview/PaymentHsmCluster_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: PaymentHsmCluster_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void paymentHsmClusterDeleteMaximumSetGen(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.paymentHsmClusters().delete("rgpaymenthsm", "phsm1", com.azure.core.util.Context.NONE);
    }
}
