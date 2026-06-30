
/**
 * Samples for Caches ResetSmbPassword.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/Caches_ResetSmbPassword.json
     */
    /**
     * Sample code: Caches_ResetSmbPassword.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void cachesResetSmbPassword(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.caches().resetSmbPassword("myResourceGroup", "account1", "pool1", "cache1",
            com.azure.core.util.Context.NONE);
    }
}
