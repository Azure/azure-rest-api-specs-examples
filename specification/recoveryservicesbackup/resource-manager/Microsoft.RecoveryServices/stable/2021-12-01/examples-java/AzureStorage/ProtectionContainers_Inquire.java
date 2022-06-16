import com.azure.core.util.Context;

/** Samples for ProtectionContainers Inquire. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/AzureStorage/ProtectionContainers_Inquire.json
     */
    /**
     * Sample code: Inquire Azure Storage Protection Containers.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void inquireAzureStorageProtectionContainers(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .protectionContainers()
            .inquireWithResponse(
                "testvault", "test-rg", "Azure", "storagecontainer;Storage;test-rg;teststorage", null, Context.NONE);
    }
}
