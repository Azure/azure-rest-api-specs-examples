/** Samples for DiagnosticsPackages Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-09-01/examples/DiagnosticsPackageGet.json
     */
    /**
     * Sample code: Get diagnostics package.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void getDiagnosticsPackage(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .diagnosticsPackages()
            .getWithResponse("rg1", "TestPacketCoreCP", "dp1", com.azure.core.util.Context.NONE);
    }
}
