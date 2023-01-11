/** Samples for Suppressions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/DeleteSuppression.json
     */
    /**
     * Sample code: DeleteSuppression.
     *
     * @param manager Entry point to AdvisorManager.
     */
    public static void deleteSuppression(com.azure.resourcemanager.advisor.AdvisorManager manager) {
        manager
            .suppressions()
            .deleteWithResponse(
                "resourceUri", "recommendationId", "suppressionName1", com.azure.core.util.Context.NONE);
    }
}
