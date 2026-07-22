
/**
 * Samples for Approval Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Approvals_Delete.json
     */
    /**
     * Sample code: Approval_Delete.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void approvalDelete(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.approvals().delete(
            "subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/TestMyRg/providers/Microsoft.Mission/enclaveconnections/TestMyEnclaveConnection",
            "TestApprovals", com.azure.core.util.Context.NONE);
    }
}
