
/**
 * Samples for PaymentHsmClusters GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01-preview/PaymentHsmCluster_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: PaymentHsmCluster_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void paymentHsmClusterGetMaximumSetGen(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.paymentHsmClusters().getByResourceGroupWithResponse("rgpaymenthsm", "phsm1",
            com.azure.core.util.Context.NONE);
    }
}
