
/**
 * Samples for PaymentHsmClusterPrivateEndpointConnections ListByPaymentHsmCluster.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-12-01-preview/PaymentHsmClusterPrivateEndpointConnection_ListByPaymentHsmCluster_MaximumSet_Gen.json
     */
    /**
     * Sample code: PaymentHsmClusterPrivateEndpointConnection_ListByPaymentHsmCluster_MaximumSet_Gen.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void paymentHsmClusterPrivateEndpointConnectionListByPaymentHsmClusterMaximumSetGen(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.paymentHsmClusterPrivateEndpointConnections().listByPaymentHsmCluster("rgpaymenthsm", "phsm1",
            com.azure.core.util.Context.NONE);
    }
}
