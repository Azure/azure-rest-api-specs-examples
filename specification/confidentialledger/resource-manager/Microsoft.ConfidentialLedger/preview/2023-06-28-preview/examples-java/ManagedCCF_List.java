
/**
 * Samples for ManagedCcf ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2023-06-28-preview/
     * examples/ManagedCCF_List.json
     */
    /**
     * Sample code: ManagedCCFList.
     * 
     * @param manager Entry point to ConfidentialLedgerManager.
     */
    public static void managedCCFList(com.azure.resourcemanager.confidentialledger.ConfidentialLedgerManager manager) {
        manager.managedCcfs().listByResourceGroup("DummyResourceGroupName", null, com.azure.core.util.Context.NONE);
    }
}
