/** Samples for AttachedDataNetworks Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-06-01/examples/AttachedDataNetworkDelete.json
     */
    /**
     * Sample code: Delete attached data network resource.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void deleteAttachedDataNetworkResource(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .attachedDataNetworks()
            .delete(
                "rg1",
                "TestPacketCoreCP",
                "TestPacketCoreDP",
                "TestAttachedDataNetwork",
                com.azure.core.util.Context.NONE);
    }
}
