
/**
 * Samples for PaymentHsmClusters ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01-preview/PaymentHsmCluster_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: PaymentHsmCluster_ListByResourceGroup_MaximumSet_Gen.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void paymentHsmClusterListByResourceGroupMaximumSetGen(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.paymentHsmClusters().listByResourceGroup("rgpaymenthsm", null, com.azure.core.util.Context.NONE);
    }
}
