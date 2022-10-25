import com.azure.core.util.Context;

/** Samples for ProtectedItems Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/preview/2022-09-01-preview/examples/Common/ProtectedItem_Delete.json
     */
    /**
     * Sample code: Delete Protection from Azure Virtual Machine.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void deleteProtectionFromAzureVirtualMachine(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .protectedItems()
            .deleteWithResponse(
                "PySDKBackupTestRsVault",
                "PythonSDKBackupTestRg",
                "Azure",
                "iaasvmcontainer;iaasvmcontainerv2;pysdktestrg;pysdktestv2vm1",
                "vm;iaasvmcontainerv2;pysdktestrg;pysdktestv2vm1",
                Context.NONE);
    }
}
