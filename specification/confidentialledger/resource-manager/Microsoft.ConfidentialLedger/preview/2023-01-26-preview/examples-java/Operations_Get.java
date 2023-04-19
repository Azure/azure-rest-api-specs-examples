/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2023-01-26-preview/examples/Operations_Get.json
     */
    /**
     * Sample code: OperationsGet.
     *
     * @param manager Entry point to ConfidentialLedgerManager.
     */
    public static void operationsGet(com.azure.resourcemanager.confidentialledger.ConfidentialLedgerManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
