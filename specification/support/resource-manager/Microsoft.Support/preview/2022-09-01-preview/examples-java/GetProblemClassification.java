/** Samples for ProblemClassifications Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/GetProblemClassification.json
     */
    /**
     * Sample code: Gets details of problemClassification for Azure service.
     *
     * @param manager Entry point to SupportManager.
     */
    public static void getsDetailsOfProblemClassificationForAzureService(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager
            .problemClassifications()
            .getWithResponse("service_guid", "problemClassification_guid", com.azure.core.util.Context.NONE);
    }
}
