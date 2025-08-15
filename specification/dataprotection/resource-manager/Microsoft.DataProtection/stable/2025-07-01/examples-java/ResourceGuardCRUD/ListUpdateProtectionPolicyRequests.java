
/**
 * Samples for ResourceGuards GetUpdateProtectionPolicyRequestsObjects.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2025-07-01/examples/
     * ResourceGuardCRUD/ListUpdateProtectionPolicyRequests.json
     */
    /**
     * Sample code: List OperationsRequestObject.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void
        listOperationsRequestObject(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.resourceGuards().getUpdateProtectionPolicyRequestsObjects("SampleResourceGroup", "swaggerExample",
            com.azure.core.util.Context.NONE);
    }
}
