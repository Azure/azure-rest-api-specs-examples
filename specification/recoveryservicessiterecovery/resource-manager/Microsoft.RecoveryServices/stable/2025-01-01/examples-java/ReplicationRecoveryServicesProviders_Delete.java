
/**
 * Samples for ReplicationRecoveryServicesProviders Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationRecoveryServicesProviders_Delete.json
     */
    /**
     * Sample code: Deletes provider from fabric. Note: Deleting provider for any fabric other than SingleHost is
     * unsupported. To maintain backward compatibility for released clients the object "deleteRspInput" is used (if the
     * object is empty we assume that it is old client and continue the old behavior).
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        deletesProviderFromFabricNoteDeletingProviderForAnyFabricOtherThanSingleHostIsUnsupportedToMaintainBackwardCompatibilityForReleasedClientsTheObjectDeleteRspInputIsUsedIfTheObjectIsEmptyWeAssumeThatItIsOldClientAndContinueTheOldBehavior(
            com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationRecoveryServicesProviders().delete("resourceGroupPS1", "vault1", "cloud1",
            "241641e6-ee7b-4ee4-8141-821fadda43fa", com.azure.core.util.Context.NONE);
    }
}
