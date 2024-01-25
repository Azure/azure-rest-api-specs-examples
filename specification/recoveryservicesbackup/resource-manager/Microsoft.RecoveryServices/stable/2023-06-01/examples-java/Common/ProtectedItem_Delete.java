
/**
 * Samples for ProtectedItems Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/
     * Common/ProtectedItem_Delete.json
     */
    /**
     * Sample code: Delete Protection from Azure Virtual Machine.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void deleteProtectionFromAzureVirtualMachine(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.protectedItems().deleteWithResponse("PySDKBackupTestRsVault", "PythonSDKBackupTestRg", "Azure",
            "iaasvmcontainer;iaasvmcontainerv2;pysdktestrg;pysdktestv2vm1",
            "vm;iaasvmcontainerv2;pysdktestrg;pysdktestv2vm1", com.azure.core.util.Context.NONE);
    }
}
