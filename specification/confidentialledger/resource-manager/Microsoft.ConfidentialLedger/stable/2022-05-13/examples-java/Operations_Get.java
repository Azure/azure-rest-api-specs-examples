import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/stable/2022-05-13/examples/Operations_Get.json
     */
    /**
     * Sample code: OperationsGet.
     *
     * @param manager Entry point to ConfidentialLedgerManager.
     */
    public static void operationsGet(com.azure.resourcemanager.confidentialledger.ConfidentialLedgerManager manager) {
        manager.operations().list(Context.NONE);
    }
}
