
/**
 * Samples for DiagnosticsPackages CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-04-01/examples/
     * DiagnosticsPackageCreate.json
     */
    /**
     * Sample code: Create diagnostics package.
     * 
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void createDiagnosticsPackage(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.diagnosticsPackages().createOrUpdate("rg1", "TestPacketCoreCP", "dp1",
            com.azure.core.util.Context.NONE);
    }
}
