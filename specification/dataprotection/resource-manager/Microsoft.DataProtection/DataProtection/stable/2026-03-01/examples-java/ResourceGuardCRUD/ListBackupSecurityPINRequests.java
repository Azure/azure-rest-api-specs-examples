
/**
 * Samples for ResourceGuards GetBackupSecurityPinRequestsObjects.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/ResourceGuardCRUD/ListBackupSecurityPINRequests.json
     */
    /**
     * Sample code: List OperationsRequestObject.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void
        listOperationsRequestObject(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.resourceGuards().getBackupSecurityPinRequestsObjects("SampleResourceGroup", "swaggerExample",
            com.azure.core.util.Context.NONE);
    }
}
