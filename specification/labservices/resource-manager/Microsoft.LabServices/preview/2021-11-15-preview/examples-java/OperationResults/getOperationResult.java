import com.azure.core.util.Context;

/** Samples for OperationResults Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/OperationResults/getOperationResult.json
     */
    /**
     * Sample code: getOperationResult.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void getOperationResult(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.operationResults().getWithResponse("a64149d8-84cb-4566-ab8e-b4ee1a074174", Context.NONE);
    }
}
