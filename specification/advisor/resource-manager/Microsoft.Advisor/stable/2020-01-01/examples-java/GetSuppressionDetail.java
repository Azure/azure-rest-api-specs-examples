/** Samples for Suppressions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/GetSuppressionDetail.json
     */
    /**
     * Sample code: GetSuppressionDetail.
     *
     * @param manager Entry point to AdvisorManager.
     */
    public static void getSuppressionDetail(com.azure.resourcemanager.advisor.AdvisorManager manager) {
        manager
            .suppressions()
            .getWithResponse("resourceUri", "recommendationId", "suppressionName1", com.azure.core.util.Context.NONE);
    }
}
