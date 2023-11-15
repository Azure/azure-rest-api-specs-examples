/** Samples for DiagnosticsPackages ListByPacketCoreControlPlane. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-09-01/examples/DiagnosticsPackageListByPacketCoreControlPlane.json
     */
    /**
     * Sample code: List diagnostics packages under a packet core control plane.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void listDiagnosticsPackagesUnderAPacketCoreControlPlane(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .diagnosticsPackages()
            .listByPacketCoreControlPlane("rg1", "TestPacketCoreCP", com.azure.core.util.Context.NONE);
    }
}
