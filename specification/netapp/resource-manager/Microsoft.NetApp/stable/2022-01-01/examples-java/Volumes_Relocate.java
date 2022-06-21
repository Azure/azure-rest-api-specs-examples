import com.azure.core.util.Context;

/** Samples for Volumes Relocate. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-01-01/examples/Volumes_Relocate.json
     */
    /**
     * Sample code: Volumes_Relocate.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesRelocate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumes().relocate("myRG", "account1", "pool1", "volume1", Context.NONE);
    }
}
