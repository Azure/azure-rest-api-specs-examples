/** Samples for SecurityPINs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/Common/BackupSecurityPin_Get.json
     */
    /**
     * Sample code: Get Vault Security Pin.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getVaultSecurityPin(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.securityPINs().getWithResponse("SwaggerTest", "SwaggerTestRg", null, com.azure.core.util.Context.NONE);
    }
}
