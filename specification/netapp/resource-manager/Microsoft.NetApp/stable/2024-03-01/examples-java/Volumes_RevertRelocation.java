
/**
 * Samples for Volumes RevertRelocation.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-03-01/examples/Volumes_RevertRelocation.json
     */
    /**
     * Sample code: Volumes_RevertRelocation.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesRevertRelocation(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumes().revertRelocation("myRG", "account1", "pool1", "volume1", com.azure.core.util.Context.NONE);
    }
}
