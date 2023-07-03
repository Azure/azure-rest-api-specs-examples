/** Samples for Racks Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/Racks_Delete.json
     */
    /**
     * Sample code: Delete rack.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void deleteRack(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.racks().delete("resourceGroupName", "rackName", com.azure.core.util.Context.NONE);
    }
}
