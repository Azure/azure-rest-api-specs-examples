/** Samples for DiagnosticsPackages Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-09-01/examples/DiagnosticsPackageDelete.json
     */
    /**
     * Sample code: Delete diagnostics package.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void deleteDiagnosticsPackage(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.diagnosticsPackages().delete("rg1", "TestPacketCoreCP", "dp1", com.azure.core.util.Context.NONE);
    }
}
