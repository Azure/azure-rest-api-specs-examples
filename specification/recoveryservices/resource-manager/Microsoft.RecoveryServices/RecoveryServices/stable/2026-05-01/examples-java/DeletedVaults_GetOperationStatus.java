
/**
 * Samples for DeletedVaults GetOperationStatus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/DeletedVaults_GetOperationStatus.json
     */
    /**
     * Sample code: Gets operation status on deleted vault.
     * 
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void
        getsOperationStatusOnDeletedVault(com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager.deletedVaults().getOperationStatusWithResponse("westus", "swaggerExample",
            "YWUzNDFkMzQtZmM5OS00MmUyLWEzNDMtZGJkMDIxZjlmZjgzOzdmYzBiMzhmLTc2NmItNDM5NS05OWQ1LTVmOGEzNzg4MWQzNA==",
            com.azure.core.util.Context.NONE);
    }
}
