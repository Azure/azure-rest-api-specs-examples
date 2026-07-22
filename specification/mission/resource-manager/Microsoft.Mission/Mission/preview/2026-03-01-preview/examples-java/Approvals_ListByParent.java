
/**
 * Samples for Approval ListByParent.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Approvals_ListByParent.json
     */
    /**
     * Sample code: Approval_ListByParent.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void approvalListByParent(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.approvals().listByParent(
            "subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/TestMyRg/providers/Microsoft.Mission/enclaveconnections/TestMyEnclaveConnection",
            com.azure.core.util.Context.NONE);
    }
}
