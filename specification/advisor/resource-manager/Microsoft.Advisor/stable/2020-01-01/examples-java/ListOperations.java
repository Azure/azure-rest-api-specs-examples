/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/ListOperations.json
     */
    /**
     * Sample code: ListRecommendations.
     *
     * @param manager Entry point to AdvisorManager.
     */
    public static void listRecommendations(com.azure.resourcemanager.advisor.AdvisorManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
