/** Samples for ManagedCcf List. */
public final class Main {
    /*
     * x-ms-original-file: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2023-01-26-preview/examples/ManagedCCF_ListBySub.json
     */
    /**
     * Sample code: ManagedCCFListBySub.
     *
     * @param manager Entry point to ConfidentialLedgerManager.
     */
    public static void managedCCFListBySub(
        com.azure.resourcemanager.confidentialledger.ConfidentialLedgerManager manager) {
        manager.managedCcfs().list(null, com.azure.core.util.Context.NONE);
    }
}
