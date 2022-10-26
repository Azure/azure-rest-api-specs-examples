import com.azure.core.util.Context;

/** Samples for ProtectionContainerOperationResults Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/preview/2022-09-01-preview/examples/AzureStorage/ProtectionContainers_Inquire_Result.json
     */
    /**
     * Sample code: Get Azure Storage Protection Container Operation Result.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getAzureStorageProtectionContainerOperationResult(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .protectionContainerOperationResults()
            .getWithResponse(
                "testvault",
                "test-rg",
                "Azure",
                "VMAppContainer;Compute;testRG;testSQL",
                "00000000-0000-0000-0000-000000000000",
                Context.NONE);
    }
}
