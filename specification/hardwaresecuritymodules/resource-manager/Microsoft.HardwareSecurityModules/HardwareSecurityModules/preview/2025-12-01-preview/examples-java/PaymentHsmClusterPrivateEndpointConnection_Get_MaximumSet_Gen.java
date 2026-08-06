
/**
 * Samples for PaymentHsmClusterPrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01-preview/PaymentHsmClusterPrivateEndpointConnection_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: PaymentHsmClusterPrivateEndpointConnection_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void paymentHsmClusterPrivateEndpointConnectionGetMaximumSetGen(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.paymentHsmClusterPrivateEndpointConnections().getWithResponse("rgpaymenthsm", "phsm1", "sample-pec",
            com.azure.core.util.Context.NONE);
    }
}
