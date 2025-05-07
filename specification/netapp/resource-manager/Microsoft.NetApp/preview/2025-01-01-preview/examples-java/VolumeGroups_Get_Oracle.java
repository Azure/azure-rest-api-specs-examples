
/**
 * Samples for VolumeGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-01-01-preview/examples/
     * VolumeGroups_Get_Oracle.json
     */
    /**
     * Sample code: VolumeGroups_Get_Oracle.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumeGroupsGetOracle(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumeGroups().getWithResponse("myRG", "account1", "group1", com.azure.core.util.Context.NONE);
    }
}
