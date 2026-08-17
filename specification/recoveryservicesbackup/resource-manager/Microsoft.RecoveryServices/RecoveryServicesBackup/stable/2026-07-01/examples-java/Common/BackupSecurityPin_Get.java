
/**
 * Samples for SecurityPINs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Common/BackupSecurityPin_Get.json
     */
    /**
     * Sample code: Get Vault Security Pin.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void
        getVaultSecurityPin(com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.securityPINs().getWithResponse("SwaggerTest", "SwaggerTestRg", null, com.azure.core.util.Context.NONE);
    }
}
