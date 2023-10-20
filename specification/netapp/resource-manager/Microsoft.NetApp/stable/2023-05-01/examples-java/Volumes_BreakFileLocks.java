import com.azure.resourcemanager.netapp.models.BreakFileLocksRequest;

/** Samples for Volumes BreakFileLocks. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-05-01/examples/Volumes_BreakFileLocks.json
     */
    /**
     * Sample code: Volumes_BreakFileLocks.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesBreakFileLocks(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager
            .volumes()
            .breakFileLocks(
                "myRG",
                "account1",
                "pool1",
                "volume1",
                new BreakFileLocksRequest().withClientIp("101.102.103.104").withConfirmRunningDisruptiveOperation(true),
                com.azure.core.util.Context.NONE);
    }
}
