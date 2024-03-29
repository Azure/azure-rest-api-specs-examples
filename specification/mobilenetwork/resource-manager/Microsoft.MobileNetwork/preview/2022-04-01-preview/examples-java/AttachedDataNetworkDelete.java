import com.azure.core.util.Context;

/** Samples for AttachedDataNetworks Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/AttachedDataNetworkDelete.json
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
            .delete("rg1", "TestPacketCoreCP", "TestPacketCoreDP", "TestAttachedDataNetwork", Context.NONE);
    }
}
