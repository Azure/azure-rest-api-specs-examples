/** Samples for ResourceGuards GetBackupSecurityPinRequestsObjects. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-05-01/examples/ResourceGuardCRUD/ListBackupSecurityPINRequests.json
     */
    /**
     * Sample code: List OperationsRequestObject.
     *
     * @param manager Entry point to DataProtectionManager.
     */
    public static void listOperationsRequestObject(
        com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager
            .resourceGuards()
            .getBackupSecurityPinRequestsObjects(
                "SampleResourceGroup", "swaggerExample", com.azure.core.util.Context.NONE);
    }
}
