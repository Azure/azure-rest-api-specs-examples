import com.azure.core.util.Context;

/** Samples for OperationResults Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/OperationResultsGet.json
     */
    /**
     * Sample code: Get operation result.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void getOperationResult(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.operationResults().getWithResponse("westus", "exampleid", Context.NONE);
    }
}
