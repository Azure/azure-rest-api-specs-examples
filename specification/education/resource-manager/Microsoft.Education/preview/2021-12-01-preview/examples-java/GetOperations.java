
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/GetOperations.
     * json
     */
    /**
     * Sample code: GetOperations.
     * 
     * @param manager Entry point to EducationManager.
     */
    public static void getOperations(com.azure.resourcemanager.education.EducationManager manager) {
        manager.operations().listWithResponse(com.azure.core.util.Context.NONE);
    }
}
