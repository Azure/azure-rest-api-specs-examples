import com.azure.core.util.Context;

/** Samples for DedicatedHsm GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/stable/2021-11-30/examples/PaymentHsm_Get_With_2018-10-31Preview_Version.json
     */
    /**
     * Sample code: Get a payment HSM with 2018-10-31Preview api version.
     *
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void getAPaymentHSMWith20181031PreviewApiVersion(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.dedicatedHsms().getByResourceGroupWithResponse("hsm-group", "hsm1", Context.NONE);
    }
}
