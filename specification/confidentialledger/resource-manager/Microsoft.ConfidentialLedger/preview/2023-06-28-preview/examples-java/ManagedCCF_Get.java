
/**
 * Samples for ManagedCcf GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2023-06-28-preview/
     * examples/ManagedCCF_Get.json
     */
    /**
     * Sample code: ManagedCCFGet.
     * 
     * @param manager Entry point to ConfidentialLedgerManager.
     */
    public static void managedCCFGet(com.azure.resourcemanager.confidentialledger.ConfidentialLedgerManager manager) {
        manager.managedCcfs().getByResourceGroupWithResponse("DummyResourceGroupName", "DummyMccfAppName",
            com.azure.core.util.Context.NONE);
    }
}
