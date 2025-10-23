
/**
 * Samples for IotDpsResource GetOperationResult.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-02-01-preview/DPSGetOperationResult.json
     */
    /**
     * Sample code: DPSGetOperationResult.
     * 
     * @param manager Entry point to IotDpsManager.
     */
    public static void
        dPSGetOperationResult(com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager.iotDpsResources().getOperationResultWithResponse("MTY5OTNmZDctODI5Yy00N2E2LTkxNDQtMDU1NGIyYzY1ZjRl",
            "myResourceGroup", "myFirstProvisioningService", "1508265712453", com.azure.core.util.Context.NONE);
    }
}
