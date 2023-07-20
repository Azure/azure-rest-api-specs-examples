import com.azure.resourcemanager.mobilenetwork.models.PacketCoreControlPlaneCollectDiagnosticsPackage;

/** Samples for PacketCoreControlPlanes CollectDiagnosticsPackage. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-06-01/examples/PacketCoreControlPlaneCollectDiagnosticsPackage.json
     */
    /**
     * Sample code: Collect diagnostics package from packet core control plane.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void collectDiagnosticsPackageFromPacketCoreControlPlane(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .packetCoreControlPlanes()
            .collectDiagnosticsPackage(
                "rg1",
                "TestPacketCoreCP",
                new PacketCoreControlPlaneCollectDiagnosticsPackage()
                    .withStorageAccountBlobUrl(
                        "https://contosoaccount.blob.core.windows.net/container/diagnosticsPackage.zip"),
                com.azure.core.util.Context.NONE);
    }
}
