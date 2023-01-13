/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/OperationsGet.json
     */
    /**
     * Sample code: OperationsGet.
     *
     * @param manager Entry point to BillingBenefitsManager.
     */
    public static void operationsGet(com.azure.resourcemanager.billingbenefits.BillingBenefitsManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
