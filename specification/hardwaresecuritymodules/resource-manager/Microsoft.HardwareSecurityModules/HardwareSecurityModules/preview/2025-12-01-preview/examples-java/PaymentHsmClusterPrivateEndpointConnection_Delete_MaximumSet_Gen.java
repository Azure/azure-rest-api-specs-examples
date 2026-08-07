
/**
 * Samples for PaymentHsmClusterPrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01-preview/PaymentHsmClusterPrivateEndpointConnection_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: PaymentHsmClusterPrivateEndpointConnection_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void paymentHsmClusterPrivateEndpointConnectionDeleteMaximumSetGen(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.paymentHsmClusterPrivateEndpointConnections().delete("rgpaymenthsm", "phsm1", "sample-pec",
            com.azure.core.util.Context.NONE);
    }
}
